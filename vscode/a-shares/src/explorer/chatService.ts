import * as vscode from 'vscode';
import globalState from '../globalState';
import Axios from 'axios';
import { LeekFundConfig } from '../shared/leekConfig';
import { StockProvider } from '../explorer/stockProvider';
import stockTrend from '../webview/stockTrend';
import homeTrend from '../webview/homeTrend';

export default class ChatViewProvider implements vscode.WebviewViewProvider {

  public static readonly viewType = 'asharesFundView.chat';

  private _view?: vscode.WebviewView;
  private interval?: NodeJS.Timeout;
  constructor(
    private readonly _extensionUri: vscode.Uri,
    private readonly _extensionContext: any,

  ) { }

  public resolveWebviewView(
    webviewView: vscode.WebviewView,
    context: vscode.WebviewViewResolveContext,
    _token: vscode.CancellationToken,
  ) {
    this._view = webviewView;

    webviewView.webview.options = {
      // Allow scripts in the webview
      enableScripts: true,

      localResourceRoots: [
        this._extensionUri,
        this._extensionContext
      ]
    };
    // const userid = globalState.userId;
    // if (userid.length === 0) {// 未登录
    //   this.updateUser();
    //   // this.interval = setInterval(this.updateUser, 3000);
    // }

    webviewView.webview.html = this._getHtmlForWebview();


    // Reset when the current panel is closed
    webviewView.onDidDispose(
      () => {
        // When the panel is closed, cancel any future updates to the webview content
        clearInterval(this.interval);
      },
      null,
      this._extensionContext
    );

    webviewView.webview.onDidReceiveMessage(data => {
      switch (data.type) {
        case 'code':
          {
            stockTrend(data.code, data.name, data.code);
            break
          }
        case 'home':
          {
            homeTrend(data.url);
            break
          }
        case 'user':
          {
            this.updateUser();
            break
          }
      }
    });
  }

  public async clearUser() {
    const deviceId = globalState.deviceId;
    const res = await Axios.post('https://www.xxxxxx.cn/shares/api/v1/weixin.clear_login', { code: deviceId });
    console.log(res)


    globalState.userId = "";
    LeekFundConfig.setConfig('a-shares.userid', globalState.userId);
    globalState.userName = "";
    LeekFundConfig.setConfig('a-shares.username', globalState.userName);
    globalState.userUrl = "";
    LeekFundConfig.setConfig('a-shares.userurl', globalState.userUrl);
    if (this._view) {
      this._view.webview.html = this._getHtmlForWebview();// 更新页面
      // this.updateUser();
    }
  }
  public async updateCode(stockProvider: StockProvider) {//
    const userId = globalState.userId;
    if (userId.length === 0) {
      vscode.window.showErrorMessage("请先扫码登录登录");
    }else if (this._view ) {
			this._view.webview.html = this._getHtmlForWebview();
		}
  }

  public async updateUser() {
    const deviceId = globalState.deviceId;
    const res = await Axios.post('https://www.xxxxxx.cn/shares/api/v1/weixin.h5_login', { code: deviceId });
    if (res.status === 200) {
      const data = res.data;
      if (data.state === true) {
        const data1 = data.data;
        if (data1.status === 1) {// 已经登陆了
          globalState.userId = data1.userId;
          LeekFundConfig.setConfig('a-shares.userid', globalState.userId);
          globalState.userName = data1.userName;
          LeekFundConfig.setConfig('a-shares.username', globalState.userName);
          globalState.userUrl = data1.avatarUrl;
          LeekFundConfig.setConfig('a-shares.userurl', globalState.userUrl);
          if (data1.rg) {
            LeekFundConfig.setConfig('a-shares.riseColor', "#ff785d");
            LeekFundConfig.setConfig('a-shares.fallColor', "#95ec69");
          } else {
            LeekFundConfig.setConfig('a-shares.riseColor', "#95ec69");
            LeekFundConfig.setConfig('a-shares.fallColor', "#ff785d");
          }
          globalState.rg = data1.rg;
          if (this._view) {
            this._view.webview.html = this._getHtmlForWebview();// 更新页面
          }
        } 
        // else {
        //   setTimeout(() => this.updateUser(), 3000)
        // }
      }
    }
    console.log('🐥>>>updateUser: ', res);
  }

  private _getHtmlForWebview() {

    const userid = globalState.userId;
    if (userid.length > 0) {// 已经登陆了
      const url = `http://localhost:8080/pc/#/chat?username=${userid}`;
      return `<!DOCTYPE html>
			<html lang="en">
			  <head>
				<meta charset="UTF-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<title>复利Chat</title>
			  </head>
			  <body>
          <div style="overflow-x:auto">
            <iframe 
            src="${url}"
            frameborder="0"
            style="width: 100%; height: 100vh"
          ></iframe>
          </div>
			  </body>
        <script>
        const vscode = acquireVsCodeApi(); // acquireVsCodeApi can only be invoked once
        window.addEventListener('message', function(event) {
            vscode.postMessage({
              type: event.data.type,
              code: event.data.code,
              name:event.data.name,
              url:event.data.url,
            })
        }, false);
        </script>
			</html>`;
    } else {
       const url = `https://www.xxxxxx.cn/webshares/#/pages/mine/login?deviceid=${globalState.deviceId}`;
      return `<!DOCTYPE html>
			<html lang="en">
			  <head>
				<meta charset="UTF-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<title>登录</title>
			  </head>
			  <body>
				<div  style="min-width: 350px; overflow-x:auto">
				  <iframe
				  src="${url}"
				  frameborder="0"
				  style="width: 100%; height: 500px"
				></iframe>
				</div>
			  </body>
        <script>
        const vscode = acquireVsCodeApi(); // acquireVsCodeApi can only be invoked once
        window.addEventListener('message', function(event) {
            console.log('======================== ', event);
            vscode.postMessage({
              type: event.data.type,
            })
        }, false);
        </script>
			</html>`;
    }

  }
}


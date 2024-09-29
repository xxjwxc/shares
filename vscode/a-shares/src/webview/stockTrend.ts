import { ViewColumn, window } from 'vscode';
import ReusedWebviewPanel from './ReusedWebviewPanel';
import globalState from '../globalState';
import { LeekFundConfig } from '../shared/leekConfig';
import { register } from 'module';



function stockTrend(code: string, name: string, stockCode: string) {
  let url = `https://www.xxxxxx.cn/pc/#/add?scode=${stockCode}`

  stockCode = stockCode.toLowerCase();
  let market = '1';
  if (stockCode.indexOf('hk') === 0) {// 港股
    market = '116';
    let mcid = market + '.' + code.substr(1);
    url = `https://quote.eastmoney.com/basic/full.html?mcid=${mcid}`
  } if (stockCode.indexOf('usr_') === 0) {// 美股
    stockCode = stockCode.replace('usr_', '');
    market = '105';
    let mcid = market + '.' + code.substr(1);
    url = `https://quote.eastmoney.com/basic/full.html?mcid=${mcid}`
  }


  const userid = globalState.userId;
  if (userid?.length > 0) {
    url += `&username=${userid}`
  } else {
    url += `&logout=true`
    window.showInformationMessage(`请先点击下方'ACCOUNT'扫码登录,效果更佳.`);
  }

  const panel = ReusedWebviewPanel.create('stockTrendWebview', name, ViewColumn.One, {
    enableScripts: true,
  });

  panel.webview.html = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>股票走势</title>
    <style>
    html.vscode-dark, body.vscode-dark, html.vscode-high-contrast, body.vscode-high-contrast {
      filter: invert(100%) hue-rotate(180deg);
    }
    </style>
  </head>
  <body>
    <div  style="overflow-x:auto">
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
        userId: event.data.userId,
				userName: event.data.userName,
				avatarUrl: event.data.avatarUrl,
				rg: event.data.rg,
      })
    }, false);
  </script>
</html>
  `;

  panel.webview.onDidReceiveMessage(data => {
    switch (data.type) {
      case 'user':
        {
          globalState.userId = data.userId;
          LeekFundConfig.setConfig('a-shares.userid', globalState.userId);
          globalState.userName = data.userName;
          LeekFundConfig.setConfig('a-shares.username', globalState.userName);
          globalState.userUrl = data.avatarUrl;
          LeekFundConfig.setConfig('a-shares.userurl', globalState.userUrl);
          if (data.rg) {
            LeekFundConfig.setConfig('a-shares.riseColor', "#ff785d");
            LeekFundConfig.setConfig('a-shares.fallColor', "#95ec69");
          } else {
            LeekFundConfig.setConfig('a-shares.riseColor', "#95ec69");
            LeekFundConfig.setConfig('a-shares.fallColor', "#ff785d");
          }
          globalState.rg = data.rg;
          console.log('🐥>>>updateUser: ', data);
          registerCallFunc();
          break;
        }
    }
  });
}

var registerCallFunc = function ()  {
  console.log('🐥>>>registerCallFunc');
}

function RegisterCallFunc(call : any) {
  registerCallFunc = call;
}



export default stockTrend;
export {RegisterCallFunc};

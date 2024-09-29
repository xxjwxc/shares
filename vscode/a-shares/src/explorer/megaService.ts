import * as vscode from 'vscode';

import  ReusedWebviewPanel  from '../webview/ReusedWebviewPanel';
import globalState from '../globalState';

export class MegaProvider implements vscode.TreeDataProvider<Node> {
  private _onDidChangeTreeData: vscode.EventEmitter<(Node | undefined)[] | undefined> = new vscode.EventEmitter<Node[] | undefined>();

  readonly onDidChangeTreeData: vscode.Event<any> = this._onDidChangeTreeData.event;


	// Keep track of any nodes we create so that we can re-use the same objects.
	private nodes: any = {};

  constructor(element: Node[]) {
    this.nodes = element;
  }

  refresh(): any {
    this._onDidChangeTreeData.fire(undefined);
  }

  public getChildren(element?: Node): Node[] {
    if (!element) {
      return this.nodes;
    } 
    return [];
  }

  getParent(element: Node): Node|undefined {
    return undefined;
  }

  getTreeItem(element: Node): vscode.TreeItem {
    const tooltip = new vscode.MarkdownString(`$(zap) Tooltip for ${element.key}`, true);
    return {
      command: {
        command: 'mega.itemClick',
        title: 'Open',
        arguments: [element],
      },
			label: { label: element.key },
			tooltip,
			resourceUri: vscode.Uri.parse(`${element.key}`),
		};
  }

   // 处理点击事件
   async onDidClickTreeItem(item: Node) {
    // 这里可以添加你想要的点击事件逻辑
    //vscode.window.showInformationMessage(`You clicked on ${item}`);
    const panel = ReusedWebviewPanel.create('stockTrendWebview', item.key, vscode.ViewColumn.One, {
      enableScripts: true,
    });
    let url = item.value
    const userid = globalState.userId;
    if (userid?.length > 0) {
      url += `&username=${userid}`
    }else {
      url += `&logout=true`
      vscode.window.showInformationMessage(`请先点击下方'ACCOUNT'扫码登录,效果更佳.`);
    }
  
   panel.webview.html = panel.webview.html = `
 <!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>${item.key}</title>
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
</html>
  
    `;
  }

}

type Node = { key: string; value: string;};
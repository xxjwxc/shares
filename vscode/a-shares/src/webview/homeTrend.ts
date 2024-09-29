import { ViewColumn, window } from 'vscode';
import ReusedWebviewPanel from './ReusedWebviewPanel';



function homeTrend(url: string) {

  const panel = ReusedWebviewPanel.create('stockTrendWebview', "QCloud", ViewColumn.One, {
    enableScripts: true,
  });

  panel.webview.html = panel.webview.html = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>QCloud</title>
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

export default homeTrend;

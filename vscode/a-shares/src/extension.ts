// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import { ConfigurationChangeEvent, ExtensionContext, TreeView, window, workspace } from 'vscode';
import { registerCommandPaletteEvent, registerViewEvent } from './registerCommand';// 注册事件
import { Telemetry } from './shared/telemetry';
import globalState from './globalState';
import StockService from './explorer/stockService';
import { StockProvider } from './explorer/stockProvider';
import  UserViewProvider  from './explorer/userService'; 
import { LeekFundConfig } from './shared/leekConfig';
import { SortType } from './shared/typed';
import { events, isStockTime } from './shared/utils';
import { MegaProvider } from './explorer/megaService';
import Axios from 'axios';
import { StatusBar } from './statusbar/statusBar';
import ChatViewProvider from './explorer/chatService';

let loopTimer: NodeJS.Timeout | undefined;
let stockTreeView: TreeView<any> | null = null;

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export async function activate(context: ExtensionContext) {
		// Use the console to output diagnostic information (console.log) and errors (console.error)
	// This line of code will only be executed once when your extension is activated
	console.log('Congratulations, your extension "a-shares" is now active!');

	globalState.context = context;

	const telemetry = new Telemetry();
	globalState.telemetry = telemetry;

	let intervalTimeConfig = LeekFundConfig.getConfig('a-shares.interval', 5000);
	let intervalTime = intervalTimeConfig;


	setGlobalVariable();

	const stockService = new StockService(context);
	const nodeStockProvider = new StockProvider(stockService);

	stockTreeView = window.createTreeView('asharesFundView.stock', {
		treeDataProvider: nodeStockProvider,
	});



  const userService = new UserViewProvider(context.extensionUri,context.subscriptions);
  context.subscriptions.push(window.registerWebviewViewProvider(UserViewProvider.viewType, userService));
  
  const chatService = new ChatViewProvider(context.extensionUri,context.subscriptions);
  context.subscriptions.push(window.registerWebviewViewProvider(ChatViewProvider.viewType, chatService));
  
  const statusBar = new StatusBar(stockService);

  var nodes  = [];
  const deviceId = globalState.deviceId;
  const res = await Axios.post('https://www.xxxxxx.cn/shares/api/v1/weixin.get_maga', { code: deviceId });
    if (res.status === 200) {
      const data = res.data;
      if (data.state === true) {
        const data1 = data.data;
        nodes = [];
        for (var i = 0; i < data1.maga.length; i++) {
          nodes.push({ key: data1.maga[i].name, value: data1.maga[i].url });
        }
      }
    }
  const megaTreeProvider = new MegaProvider(nodes)
  context.subscriptions.push(window.registerTreeDataProvider('asharesFundView.mega', megaTreeProvider));


  
  console.log('🐥>>>deviceid: ', globalState.deviceId);// VSCODE_lx0gd8fd_tqupgbgod

  // fix when TreeView collapse https://github.com/giscafer/a-shares/issues/31
	const manualRequest = () => {
		stockService.getData(LeekFundConfig.getConfig('a-shares.stocks'), SortType.NORMAL);
	};

	manualRequest();

  // loop
  const loopCallback = () => {
    if (isStockTime()) {
      // 重置定时器
      if (intervalTime !== intervalTimeConfig) {
        intervalTime = intervalTimeConfig;
        setIntervalTime();
        return;
      }

      if (stockTreeView?.visible ) {
        nodeStockProvider.refresh();
      } else {
        manualRequest();
      }
    } else {
      console.log('StockMarket Closed! Polling closed!');
      // 闭市时增加轮询间隔时长
      if (intervalTime === intervalTimeConfig) {
        intervalTime = intervalTimeConfig * 100;
        setIntervalTime();
      }
    }
  };

  const setIntervalTime = () => {
    // prevent qps
    if (intervalTime < 3000) {
      intervalTime = 3000;
    }
    if (loopTimer) {
      clearInterval(loopTimer);
      loopTimer = undefined;
    }

    loopTimer = setInterval(loopCallback, intervalTime);

  };

  setIntervalTime();

  workspace.onDidChangeConfiguration((e: ConfigurationChangeEvent) => {
    console.log('🐥>>>Configuration changed', e);
    intervalTimeConfig = LeekFundConfig.getConfig('a-shares.interval');
    setIntervalTime();
    setGlobalVariable();
    nodeStockProvider.refresh();
    events.emit('onDidChangeConfiguration');
  });

	  // register event
	  registerViewEvent(
		context,
		stockService,
    userService,
    chatService,
		nodeStockProvider,
    megaTreeProvider,
	  );
	
	  // register command
	   registerCommandPaletteEvent(context, statusBar);
	
	  // Telemetry Event
	  telemetry.sendEvent('activate');
	  
}

function setGlobalVariable() {
	globalState.iconType = LeekFundConfig.getConfig('a-shares.iconType') || 'arrow';
  globalState.rg = LeekFundConfig.getConfig('a-shares.riseColor') == "#ff785d";
	globalState.labelFormat = LeekFundConfig.getConfig('a-shares.labelFormat');
  globalState.deviceId = LeekFundConfig.getConfig('a-shares.deviceid', "");
  if (!globalState.deviceId) {
    globalState.deviceId = 'VSCODE_' + Date.now().toString(36) + '_' + Math.random().toString(36).substr(2,9)
    LeekFundConfig.setConfig('a-shares.deviceid', globalState.deviceId);
  }
  globalState.userId = LeekFundConfig.getConfig('a-shares.userid', "");
  globalState.userName = LeekFundConfig.getConfig('a-shares.username', "");
  globalState.userUrl = LeekFundConfig.getConfig('a-shares.userurl', "");
}

// This method is called when your extension is deactivated
export function deactivate() {
	console.log('🐥deactivate');
	if (loopTimer) {
	  clearInterval(loopTimer);
	  loopTimer = undefined;
	}
}


type Node = { key: string; value: string;};
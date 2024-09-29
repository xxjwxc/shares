import { commands, ExtensionContext, window } from 'vscode';
// import { FundProvider } from './explorer/fundProvider';
import { StockProvider } from './explorer/stockProvider';
import StockService from './explorer/stockService';
// import globalState from './globalState';
import { LeekFundConfig } from './shared/leekConfig';
// import { LeekTreeItem } from './shared/leekTreeItem';
// import checkForUpdate from './shared/update';
// import { colorOptionList, randomColor } from './shared/utils';
// import donate from './webview/donate';
// import leekCenterView from './webview/leekCenterView';

import stockTrend, { RegisterCallFunc } from './webview/stockTrend';
import UserViewProvider from './explorer/userService';
import { MegaProvider } from './explorer/megaService';
import globalState from './globalState';
import Axios from 'axios'; 
import { StatusBar } from './statusbar/statusBar';
import { LeekTreeItem } from './shared/leekTreeItem';
import ChatViewProvider from './explorer/chatService';
// import tucaoForum from './webview/tucaoForum';
// import { StatusBar } from './statusbar/statusBar';
let codes = "";
let yjzd = false;


export function registerViewEvent(
  context: ExtensionContext,
  stockService: StockService,
  userService: UserViewProvider,
  chatService: ChatViewProvider,
  stockProvider: StockProvider,
  megaTreeProvider: MegaProvider,
) {

    // Stock operation
    commands.registerCommand('a-shares.clearStock', () => {
      LeekFundConfig.clearStock('sh000001', () => {
        stockProvider.refresh();
       
      });
    });

  // Stock operation
  commands.registerCommand('a-shares.refreshStock', () => {
    stockProvider.refresh();
    const handler = window.setStatusBarMessage(`股票数据已刷新`);
    setTimeout(() => {
      handler.dispose();
    }, 1000);
  });

  
  // 一键诊断
  commands.registerCommand('a-shares.yjzd', async () => {
      const userId = globalState.userId;
      if(userId.length === 0) {
        window.setStatusBarMessage(`股票数据已刷新`);
        return;
      }
      yjzd = !yjzd;
      if (yjzd) {
        codes = LeekFundConfig.getConfig('a-shares.stocks');
        const res = await Axios.post('https://www.xxxxxx.cn/shares/api/v1/shares.yjzd_code',{codes: codes},
          {
            headers : {
              'Content-Type': 'application/json',
              'user-token': userId
              }
          }
        );
        if (res.status === 200) {
          const data = res.data;
          if (data.state === true) {
            const data1 = data.data;
            let coodes = ""
            for(let i=0;i<data1.codes.length;i++) {
              coodes+=data1.codes[i] + ","
            }
    
            LeekFundConfig.replaceStockCfg(coodes, () => {
              stockProvider.refresh();
              });
              const handler = window.setStatusBarMessage(`一键诊断已完成`);
              setTimeout(() => {
                handler.dispose();
              }, 1000);
          }else {
            const handler = window.setStatusBarMessage(data.error);
            setTimeout(() => {
              handler.dispose();
            }, 1000);
          }
        }

      }else {
        let coodes = ""
        for(let i=0;i<codes.length;i++) {
          coodes+=codes[i] + ","
        }
        LeekFundConfig.replaceStockCfg(coodes, () => {
          stockProvider.refresh();
          });
      }
  });

  // 删除股票
  commands.registerCommand('a-shares.deleteStock', (target) => {
    LeekFundConfig.removeStockCfg(target.id, () => {
      stockProvider.refresh();
    });
  });
  // 删除股票
   commands.registerCommand('a-shares.deleteStockOnline', (target) => {
    LeekFundConfig.removeStockCfg(target.id, () => {
      stockProvider.refresh();
      const userId = globalState.userId;
      if(userId.length === 0) {
        window.showErrorMessage("请先扫码登录登录");
        return;
      }
      Axios.post('https://www.xxxxxx.cn/shares/api/v1/shares.delete_my_code',{		code: target.id},
        {
          headers : {
            'Content-Type': 'application/json',
            'user-token': userId
            }
        }
      );
    });
  });
  commands.registerCommand('a-shares.addStockToBar', (target) => {
    LeekFundConfig.addStockToBarCfg(target.id, () => {
      stockProvider.refresh();
    });
  });
  commands.registerCommand('a-shares.leekCenterView', () => {
    if (stockService.stockList.length === 0 ) {
      window.showWarningMessage('数据刷新中，请稍候！');
      return;
    }
    //leekCenterView(stockService);
  });

  commands.registerCommand('a-shares.addStock', () => {
    // vscode QuickPick 不支持动态查询，只能用此方式解决
    // https://github.com/microsoft/vscode/issues/23633
    const qp = window.createQuickPick();
    qp.items = [{ label: '请输入关键词查询，如：0000001 或 上证指数 或者拼音首字母' }];
    let code: string | undefined;
    let timer: NodeJS.Timeout | undefined;
    qp.onDidChangeValue((value) => {
      qp.busy = true;
      if (timer) {
        clearTimeout(timer);
        timer = undefined;
      }
      timer = setTimeout(async () => {
        const res = await stockService.getStockSuggestList(value);
        qp.items = res;
        qp.busy = false;
      }, 100); // 简单防抖
    });
    qp.onDidChangeSelection((e) => {
      if (e[0].description) {
        code = e[0].label && e[0].label.split(' | ')[0];
      }
    });
    qp.show();
    qp.onDidAccept(() => {
      if (!code) {
        return;
      }
      // 存储到配置的时候是接口的参数格式，接口请求时不需要再转换
      const newCode = code.replace('gb', 'gb_').replace('us', 'usr_');
      LeekFundConfig.updateStockCfg(newCode, () => {
        stockProvider.refresh();
      });


      qp.hide();
      qp.dispose();
    });
  });
  
  commands.registerCommand('a-shares.sortStock', () => {
    stockProvider.changeOrder();
    stockProvider.refresh();
  });

  /**
   * WebView
   */
  // 股票点击
  context.subscriptions.push(
    commands.registerCommand('a-shares.stockItemClick', (code, name, text, stockCode) =>
      stockTrend(code, name, stockCode)
    )
  );
  // 状态栏
  context.subscriptions.push(
    commands.registerCommand('a-shares.changeStatusBarItem', (stockId) => {
      const stockList = stockService.stockList;
      const stockNameList = stockList
        .filter((stock) => stock.id !== stockId)
        .map((item: LeekTreeItem) => {
          return {
            label: `${item.info.name}`,
            description: `${item.info.code}`,
          };
        });
      stockNameList.unshift({
        label: `删除`,
        description: `-1`,
      });
      window
        .showQuickPick(stockNameList, {
          placeHolder: '更换状态栏个股',
        })
        .then((res) => {
          if (!res) return;
          const statusBarStocks = LeekFundConfig.getConfig('a-shares.statusBarStock');
          const newCfg = [...statusBarStocks];
          const newStockId = res.description;
          const index = newCfg.indexOf(stockId);
          if (newStockId === '-1') {
            if (index > -1) {
              newCfg.splice(index, 1);
            }
          } else {
            if (statusBarStocks.includes(newStockId)) {
              window.showWarningMessage(`「${res.label}」已在状态栏`);
              return;
            }
            if (index > -1) {
              newCfg[index] = res.description;
            }
          }
          LeekFundConfig.updateStatusBarStockCfg(newCfg, () => {
            const handler = window.setStatusBarMessage(`下次数据刷新见效`);
            setTimeout(() => {
              handler.dispose();
            }, 1500);
          });
        });
    })
  );

  // 股票置顶
  commands.registerCommand('a-shares.setStockTop', (target) => {
    LeekFundConfig.setStockTopCfg(target.id, () => {
      // fundProvider.refresh();
      //stockProvider.refresh()
    });
  });
  // 股票上移
  commands.registerCommand('a-shares.setStockUp', (target) => {
    LeekFundConfig.setStockUpCfg(target.id, () => {
      // fundProvider.refresh();
     //  stockProvider.refresh()
    });
  });
  // 股票下移
  commands.registerCommand('a-shares.setStockDown', (target) => {
    LeekFundConfig.setStockDownCfg(target.id, () => {
      // fundProvider.refresh();
      //stockProvider.refresh()
    });
  });

  context.subscriptions.push(
		commands.registerCommand('a-shares.clearUser', () => {
			userService.clearUser();
      chatService.clearUser();
		}));

    context.subscriptions.push(
    commands.registerCommand('a-shares.updateCode', () => {
			userService.updateCode(stockProvider);
      chatService.updateCode(stockProvider);
		}));
    RegisterCallFunc(()=>{
      userService.updateCode(stockProvider);
      chatService.updateCode(stockProvider);
    });
    // 监听节点被点击的事件
    context.subscriptions.push(commands.registerCommand('mega.itemClick', (item) => {
      megaTreeProvider.onDidClickTreeItem(item);
    }));
}

export function registerCommandPaletteEvent(context: ExtensionContext, statusbar: StatusBar) {
  context.subscriptions.push(
    commands.registerCommand('a-shares.toggleStatusBarVisibility', () => {
      statusbar.toggleVisibility();
    })
  );
}

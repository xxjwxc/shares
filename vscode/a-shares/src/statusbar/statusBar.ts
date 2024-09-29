import { StatusBarAlignment, StatusBarItem, window } from 'vscode';
import StockService from '../explorer/stockService';
import globalState from '../globalState';
import { DEFAULT_LABEL_FORMAT } from '../shared/constant';
import { LeekFundConfig } from '../shared/leekConfig';
import { LeekTreeItem } from '../shared/leekTreeItem';
import { events, formatLabelString } from '../shared/utils';

export class StatusBar {
  private stockService: StockService;
  private statusBarList: StatusBarItem[] = [];
  private statusBarItemLabelFormat: string = '';
  constructor(stockService: StockService) {
    this.stockService = stockService;
    this.statusBarList = [];
    this.refreshStockStatusBar();
    this.bindEvents();
    /* events.on('updateConfig:a-shares.statusBarStock',()=>{

    }) */
  }

  get riseColor(): string {
    return LeekFundConfig.getConfig('a-shares.riseColor');
  }

  get fallColor(): string {
    return LeekFundConfig.getConfig('a-shares.fallColor');
  }

  /** 隐藏股市状态栏 */
  get hideStatusBarStock(): boolean {
    return LeekFundConfig.getConfig('a-shares.hideStatusBarStock');
  }
  /** 隐藏状态栏 */
  get hideStatusBar(): boolean {
    return LeekFundConfig.getConfig('a-shares.hideStatusBar');
  }

  bindEvents() {
    events.on('stockListUpdate', () => {
      this.refreshStockStatusBar();
    });
  }

  refresh() {
    this.refreshStockStatusBar();
  }

  toggleVisibility() {
    LeekFundConfig.setConfig('a-shares.hideStatusBar', !this.hideStatusBar);
    this.refresh();
  }

  refreshStockStatusBar() {
    if (this.hideStatusBar || this.hideStatusBarStock || !this.stockService.stockList.length) {
      if (this.statusBarList.length) {
        this.statusBarList.forEach((bar) => bar.dispose());
        this.statusBarList = [];
      }
      return;
    }

    let sz: LeekTreeItem | null = null;
    const statusBarStocks = LeekFundConfig.getConfig('a-shares.statusBarStock');
    const barStockList: Array<LeekTreeItem> = new Array(statusBarStocks.length);

    this.statusBarItemLabelFormat =
      globalState.labelFormat?.['statusBarLabelFormat'] ??
      DEFAULT_LABEL_FORMAT.statusBarLabelFormat;

    this.stockService.stockList.forEach((stockItem) => {
      const { code } = stockItem.info;
      if (code === 'sh000001') {
        sz = stockItem;
      }
      if (statusBarStocks.includes(code)) {
        // barStockList.push(stockItem);
        barStockList[statusBarStocks.indexOf(code)] = stockItem;
      }
    });

    if (!barStockList.length) {
      barStockList.push(sz || this.stockService.stockList[0]);
    }

    let count = barStockList.length - this.statusBarList.length;
    if (count > 0) {
      while (--count >= 0) {
        const stockBarItem = window.createStatusBarItem(StatusBarAlignment.Left, 3);
        this.statusBarList.push(stockBarItem);
      }
    } else if (count < 0) {
      let num = Math.abs(count);
      while (--num >= 0) {
        const bar = this.statusBarList.pop();
        bar?.hide();
        bar?.dispose();
      }
    }
    barStockList.forEach((stock, index) => {
      this.updateBarInfo(this.statusBarList[index], stock);
    });
  }

  updateBarInfo(stockBarItem: StatusBarItem, item: LeekTreeItem | null) {
    if (!item) return;
    const { code, percent, open,price, yestclose, high, low, updown, amount } = item.info;
    const deLow = percent.indexOf('-') === -1;
   stockBarItem.text = `「${this.stockService.showLabel ? item.info.name : item.id}」${price}(${percent}%)`;
    // stockBarItem.text = formatLabelString(this.statusBarItemLabelFormat, {
    //   ...item.info,
    //   percent: `${percent}%`,
    //   icon: deLow ? '📈' : '📉',
    // });

    stockBarItem.tooltip = `「今日行情」 ${
      item.info?.name ?? '今日行情'
    }（${code}）\n涨跌：${updown}   百分：${percent}%\n最高：${high}   最低：${low}\n今开：${open}   昨收：${yestclose}\n成交额：${amount}\n更新时间：${
      item.info?.time
    }`;
    stockBarItem.color = deLow ? this.riseColor : this.fallColor;
    stockBarItem.command = {
      title: 'Change stock',
      command: 'a-shares.changeStatusBarItem',
      arguments: [item.id],
    };

    stockBarItem.show();
    return stockBarItem;
  }

}

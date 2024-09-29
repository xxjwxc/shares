/*--------------------------------------------------------------
 *  Copyright (c) Nickbing Lao<giscafer@outlook.com>. All rights reserved.
 *  Licensed under the MIT License.
 *  Github: https://github.com/giscafer
 *-------------------------------------------------------------*/

import { window, workspace } from 'vscode';
import globalState from '../globalState';
import { clean, uniq, events } from './utils';

export class BaseConfig {
  static getConfig(key: string, defaultValue?: any): any {
    const config = workspace.getConfiguration();
    const value = config.get(key);
    return value === undefined ? defaultValue : value;
  }

  static setConfig(cfgKey: string, cfgValue: Array<any> | string | number | Object) {
    events.emit('updateConfig:' + cfgKey, cfgValue);
    const config = workspace.getConfiguration();
    return config.update(cfgKey, cfgValue, true);
  }

  static updateConfig(cfgKey: string, codes: Array<any>) {
    const config = workspace.getConfiguration();
    const sourceCfg = config.get(cfgKey, []);

    const newCfg = sourceCfg.filter((item) => !codes.includes(item));

    const updatedCfg = [...newCfg, ...codes];
    let newCodes = clean(updatedCfg);
    newCodes = uniq(newCodes);
    return config.update(cfgKey, newCodes, true);
  }

  static clearConfig(cfgKey: string, codes: Array<any>) {
    const config = workspace.getConfiguration();

    return config.update(cfgKey, codes, true);
  }

  static removeConfig(cfgKey: string, code: string) {
    const config = workspace.getConfiguration();
    const sourceCfg = config.get(cfgKey, []);
    const newCfg = sourceCfg.filter((item) => item !== code);
    if(sourceCfg.length === newCfg.length){
      window.showInformationMessage(`删除期货不成功。请 [点击此处](https://github.com/LeekHub/a-shares/issues/281) 查看期货相关问题`);
    }
    return config.update(cfgKey, newCfg, true);
  }
}

export class LeekFundConfig extends BaseConfig {
  constructor() {
    super();
  }

    // Stock Begin
    static clearStock(codes: string, cb?: Function) {
      this.clearConfig('a-shares.stocks',codes.split(',')).then(() => {
        window.showInformationMessage(`数据已清空.`);
        if (cb && typeof cb === 'function') {
          cb(codes);
        }
      });
    }

  // Stock Begin
  static updateStockCfg(codes: string, cb?: Function) {
    this.updateConfig('a-shares.stocks', codes.split(',')).then(() => {
      window.showInformationMessage(`Stock Successfully add.`);
      if (cb && typeof cb === 'function') {
        cb(codes);
      }
    });
  }

    // Stock Begin
    static replaceStockCfg(codes: string, cb?: Function) {
      this.clearConfig('a-shares.stocks', codes.split(',')).then(() => {
        window.showInformationMessage(`Stock Successfully add.`);
        if (cb && typeof cb === 'function') {
          cb(codes);
        }
      });
    }

  static removeStockCfg(code: string, cb?: Function) {
    this.removeConfig('a-shares.stocks', code).then(() => {
      window.showInformationMessage(`Stock Successfully delete.`);
      if (cb && typeof cb === 'function') {
        cb(code);
      }
    });
  }

  static addStockToBarCfg(code: string, cb?: Function) {
    const addStockToBar = () => {
      let configArr: string[] = this.getConfig('a-shares.statusBarStock');
      if (configArr.length >= 4) {
        window.showInformationMessage(`StatusBar Exceeding Length.`);
        if (cb && typeof cb === 'function') {
          cb(code);
        }
      } else if (configArr.includes(code)) {
        window.showInformationMessage(`StatusBar Already Have.`);
        if (cb && typeof cb === 'function') {
          cb(code);
        }
      } else {
        configArr.push(code);
        this.setConfig('a-shares.statusBarStock', configArr).then(() => {
          window.showInformationMessage(`Stock Successfully add to statusBar.`);
          if (cb && typeof cb === 'function') {
            cb(code);
          }
        });
      }
    };

    if (this.getConfig('a-shares.hideStatusBarStock')) {
      this.setConfig('a-shares.hideStatusBarStock', false).then(() => {
        addStockToBar();
      });
    } else {
      addStockToBar();
    }
  }

  static setStockTopCfg(code: string, cb?: Function) {
    let configArr: string[] = this.getConfig('a-shares.stocks');

    configArr = [code, ...configArr.filter((item) => item !== code)];

    this.setConfig('a-shares.stocks', configArr).then(() => {
      window.showInformationMessage(`Stock successfully set to top.`);
      if (cb && typeof cb === 'function') {
        cb(code);
      }
    });
  }

  static setStockUpCfg(code: string, cb?: Function) {
    const callback = () => {
      window.showInformationMessage(`Stock successfully move up.`);
      if (cb && typeof cb === 'function') {
        cb(code);
      }
    };

    let configArr: string[] = this.getConfig('a-shares.stocks');
    const currentIndex = configArr.indexOf(code);
    let previousIndex = currentIndex - 1;
    // 找到前一个同市场的股票
    for (let index = currentIndex - 1; index >= 0; index--) {
      const previousCode = configArr[index];
      if (/^(sh|sz|bj)/.test(code) && /^(sh|sz|bj)/.test(previousCode)) {
        previousIndex = index;
        break;
      }
      if (/^(hk)/.test(code) && /^(hk)/.test(previousCode)) {
        previousIndex = index;
        break;
      }
      if (/^(usr_)/.test(code) && /^(usr_)/.test(previousCode)) {
        previousIndex = index;
        break;
      }
      if (/^(nf_)/.test(code) && /^(nf_)/.test(previousCode)) {
        previousIndex = index;
        break;
      }
      if (/^(hf_)/.test(code) && /^(hf_)/.test(previousCode)) {
        previousIndex = index;
        break;
      }
    }
    if (previousIndex < 0) {
      callback();
    } else {
      // 交换位置
      configArr[currentIndex] = configArr.splice(previousIndex, 1, configArr[currentIndex])[0];
      this.setConfig('a-shares.stocks', configArr).then(() => {
        callback();
      });
    }
  }

  static setStockDownCfg(code: string, cb?: Function) {
    const callback = () => {
      window.showInformationMessage(`Stock successfully move down.`);
      if (cb && typeof cb === 'function') {
        cb(code);
      }
    };

    let configArr: string[] = this.getConfig('a-shares.stocks');
    const currentIndex = configArr.indexOf(code);
    let nextIndex = currentIndex + 1;
    //找到后一个同市场的股票
    for (let index = currentIndex + 1; index < configArr.length; index++) {
      const nextCode = configArr[index];
      if (/^(sh|sz|bj)/.test(code) && /^(sh|sz|bj)/.test(nextCode)) {
        nextIndex = index;
        break;
      }
      if (/^(hk)/.test(code) && /^(hk)/.test(nextCode)) {
        nextIndex = index;
        break;
      }
      if (/^(usr_)/.test(code) && /^(usr_)/.test(nextCode)) {
        nextIndex = index;
        break;
      }
      if (/^(nf_)/.test(code) && /^(nf_)/.test(nextCode)) {
        nextIndex = index;
        break;
      }
      if (/^(hf_)/.test(code) && /^(hf_)/.test(nextCode)) {
        nextIndex = index;
        break;
      }
    }
    if (nextIndex >= configArr.length) {
      callback();
    } else {
      // 交换位置
      configArr[currentIndex] = configArr.splice(nextIndex, 1, configArr[currentIndex])[0];
      this.setConfig('a-shares.stocks', configArr).then(() => {
        callback();
      });
    }
  }

  // Stock End

  // Binance Begin
  static updateBinanceCfg(codes: string, cb?: Function) {
    this.updateConfig('a-shares.binance', codes.split(',')).then(() => {
      window.showInformationMessage(`Pair Successfully add.`);
      if (cb && typeof cb === 'function') {
        cb(codes);
      }
    });
  }
  static removeBinanceCfg(code: string, cb?: Function) {
    this.removeConfig('a-shares.binance', code).then(() => {
      window.showInformationMessage(`Pair Successfully delete.`);
      if (cb && typeof cb === 'function') {
        cb(code);
      }
    });
  }
  static setBinanceTopCfg(code: string, cb?: Function) {
    let configArr: string[] = this.getConfig('a-shares.binance');
    configArr = [code, ...configArr.filter((item) => item !== code)];
    this.setConfig('a-shares.binance', configArr).then(() => {
      window.showInformationMessage(`Pair successfully set to top.`);
      if (cb && typeof cb === 'function') {
        cb(code);
      }
    });
  }
  // Binance end

  // StatusBar Begin
  static updateStatusBarStockCfg(codes: Array<string>, cb?: Function) {
    const updateStatusBarStock = () => {
      this.setConfig('a-shares.statusBarStock', codes).then(() => {
        window.showInformationMessage(`Status Bar Stock Successfully update.`);
        if (cb && typeof cb === 'function') {
          cb(codes);
        }
      });
    };

    if (codes.length) {
      if (this.getConfig('a-shares.hideStatusBarStock')) {
        this.setConfig('a-shares.hideStatusBarStock', false).then(() => {
          updateStatusBarStock();
        });
      } else {
        updateStatusBarStock();
      }
    } else {
      if (!this.getConfig('a-shares.hideStatusBarStock')) {
        this.setConfig('a-shares.hideStatusBarStock', true).then(() => {
          updateStatusBarStock();
        });
      } else {
        updateStatusBarStock();
      }
    }
  }
  // StatusBar End
}

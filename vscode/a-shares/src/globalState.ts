import { ExtensionContext } from 'vscode';
import { DEFAULT_LABEL_FORMAT } from './shared/constant';
import { Telemetry } from './shared/telemetry';

let deviceId : string = '';
let userId : string = '';
let userName : string = '';
let userUrl : string = '';

let context: ExtensionContext = undefined as unknown as ExtensionContext;

let telemetry: Telemetry | any = null;
let rg = true;

let iconType = 'arrow';

let stocksRemind: Record<string, any> = {};
let newsIntervalTime = 20000; // 新闻刷新频率（毫秒）
let newsIntervalTimer: NodeJS.Timer | any = null; // 计算器控制
let labelFormat = DEFAULT_LABEL_FORMAT;


let aStockCount = 0;
let usStockCount = 0;
let hkStockCount = 0;
let noDataStockCount = 0;
let isHolidayChina = false; // 初始化状态，默认是false，免得API有问题，就当它不是好了，可以继续运行

let showStockErrorInfo = true; // 控制只显示一次错误弹窗（临时处理）

let isDevelopment = false; // 是否开发环境


let stockPrice = {}; // 缓存数据
let stockPriceCacheDate = '2020-10-30';
export default {
  context,
  telemetry,
  iconType,
  deviceId,
  userId,
  userName,
  userUrl,
  newsIntervalTime,
  newsIntervalTimer,
  aStockCount,
  usStockCount,
  hkStockCount,
  noDataStockCount,
  /**
   * 当天是否中国节假日（在插件启动时获取）
   */
  isHolidayChina,
  stocksRemind,
  labelFormat,
  showStockErrorInfo,
  isDevelopment,

  stockPrice,
  stockPriceCacheDate,
  rg,

};

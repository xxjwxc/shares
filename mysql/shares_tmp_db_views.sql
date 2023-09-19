/*
 Navicat Premium Data Transfer

 Source Server         : 朝泥沟(175.24.103.30)
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : 175.24.103.30:3306
 Source Schema         : shares_tmp_db

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 23/05/2022 17:33:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- View structure for analy_cd_view
-- ----------------------------
DROP VIEW IF EXISTS `analy_cd_view`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `analy_cd_view` AS select `a`.`id` AS `id`,`a`.`code` AS `code`,`a`.`name` AS `name`,`a`.`day0` AS `day0`,`a`.`day_str` AS `day_str`,`a`.`price` AS `price`,`a`.`new_price` AS `new_price`,`a`.`percent` AS `percent`,`a`.`turnover_rate` AS `turnover_rate`,`a`.`score` AS `score`,`a`.`doc` AS `doc`,`a`.`created_at` AS `created_at`,`c`.`hy_name` AS `hy_name` from ((`analy_cd_tbl` `a` join `good_shares_tbl` `b` on((`b`.`code` = `a`.`code`))) join `shares_info_tbl` `c` on((`c`.`code` = `a`.`code`)));

-- ----------------------------
-- View structure for analy_fl_view
-- ----------------------------
DROP VIEW IF EXISTS `analy_fl_view`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `analy_fl_view` AS select `a`.`id` AS `id`,`a`.`code` AS `code`,`a`.`name` AS `name`,`c`.`name` AS `up_name`,`a`.`day` AS `day`,`a`.`day_str` AS `day_str`,`a`.`price` AS `price`,`a`.`new_price` AS `new_price`,`a`.`percent` AS `percent`,`a`.`turnover_rate` AS `turnover_rate`,`a`.`score` AS `score`,`a`.`created_at` AS `created_at`,`c`.`num` AS `num`,`c`.`up` AS `up`,`b`.`hy_name` AS `hy_name` from ((`analy_fl_tbl` `a` join `shares_info_tbl` `b` on((`b`.`code` = `a`.`code`))) join `analy_hy_tbl` `c` on((`b`.`hy_name` like concat('%',`c`.`name`,'%'))));

-- ----------------------------
-- View structure for analy_fl_view_old
-- ----------------------------
DROP VIEW IF EXISTS `analy_fl_view_old`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `analy_fl_view_old` AS select `a`.`id` AS `id`,`a`.`code` AS `code`,`a`.`name` AS `name`,`c`.`name` AS `up_name`,`a`.`day` AS `day`,`a`.`day_str` AS `day_str`,`a`.`price` AS `price`,`a`.`new_price` AS `new_price`,`a`.`percent` AS `percent`,`a`.`turnover_rate` AS `turnover_rate`,`a`.`score` AS `score`,`a`.`created_at` AS `created_at`,`c`.`num` AS `num`,`c`.`up` AS `up`,`b`.`hy_name` AS `hy_name` from ((`analy_fl_tbl` `a` join `shares_info_tbl` `b` on((`b`.`code` = `a`.`code`))) join `hy_up_tbl` `c` on((`b`.`hy_name` like concat('%',`c`.`name`,'%'))));

-- ----------------------------
-- View structure for analy_hy_view
-- ----------------------------
DROP VIEW IF EXISTS `analy_hy_view`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `analy_hy_view` AS select `a`.`hy_code` AS `hy_code`,`a`.`hy_name` AS `hy_name`,(select sum(`hy_daily_tbl`.`zljlr`) from `hy_daily_tbl` where ((`hy_daily_tbl`.`hy_code` = `a`.`hy_code`) and (`hy_daily_tbl`.`day0` >= (select distinct `hy_daily_tbl`.`day0` from `hy_daily_tbl` order by `hy_daily_tbl`.`day0` desc limit 0,1)))) AS `price1`,(select sum(`hy_daily_tbl`.`zljlr`) from `hy_daily_tbl` where ((`hy_daily_tbl`.`hy_code` = `a`.`hy_code`) and (`hy_daily_tbl`.`day0` >= (select distinct `hy_daily_tbl`.`day0` from `hy_daily_tbl` order by `hy_daily_tbl`.`day0` desc limit 4,1)))) AS `price5`,(select sum(`hy_daily_tbl`.`zljlr`) from `hy_daily_tbl` where ((`hy_daily_tbl`.`hy_code` = `a`.`hy_code`) and (`hy_daily_tbl`.`day0` >= (select distinct `hy_daily_tbl`.`day0` from `hy_daily_tbl` order by `hy_daily_tbl`.`day0` desc limit 9,1)))) AS `price10`,(select sum(`hy_daily_tbl`.`zljlr`) from `hy_daily_tbl` where ((`hy_daily_tbl`.`hy_code` = `a`.`hy_code`) and (`hy_daily_tbl`.`day0` >= (select distinct `hy_daily_tbl`.`day0` from `hy_daily_tbl` order by `hy_daily_tbl`.`day0` desc limit 19,1)))) AS `price20`,`a`.`day0` AS `day0`,`a`.`day0_str` AS `day0_str` from `hy_daily_tbl` `a` where (`a`.`day0` = (select max(`hy_daily_tbl`.`day0`) from `hy_daily_tbl`)) order by `price1` desc,`price5` desc,`price10` desc,`price20` desc;

-- ----------------------------
-- View structure for analy_tdx_after_view
-- ----------------------------
DROP VIEW IF EXISTS `analy_tdx_after_view`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `analy_tdx_after_view` AS select `a`.`id` AS `id`,`a`.`code` AS `code`,`a`.`name` AS `name`,`a`.`day0_str` AS `day0_str`,`a`.`price` AS `price`,`a`.`new_price` AS `new_price`,`c`.`hy_name` AS `hy_name`,`a`.`day0` AS `day0`,`a`.`min` AS `min`,`a`.`max` AS `max`,`a`.`percent` AS `percent`,`a`.`fyyx` AS `fyyx`,`a`.`zlzxh` AS `zlzxh`,`a`.`fx` AS `fx`,`a`.`smx` AS `smx`,`a`.`dwzq` AS `dwzq`,`a`.`ksls` AS `ksls`,`a`.`jdz` AS `jdz`,`a`.`doc` AS `doc`,`a`.`count` AS `count`,`a`.`turnover_rate` AS `turnover_rate`,`a`.`created_at` AS `created_at` from (((`analy_tdx_after_tbl` `a` join `good_shares_tbl` `b` on((`b`.`code` = `a`.`code`))) join `shares_info_tbl` `c` on((`c`.`code` = `a`.`code`))) join `analy_tdx_tbl` `d` on((`a`.`code` = `d`.`code`))) order by `a`.`day0` desc;

-- ----------------------------
-- View structure for analy_tdx_view
-- ----------------------------
DROP VIEW IF EXISTS `analy_tdx_view`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `analy_tdx_view` AS select `a`.`id` AS `id`,`a`.`code` AS `code`,`a`.`name` AS `name`,`a`.`day0_str` AS `day0_str`,`a`.`price` AS `price`,`a`.`new_price` AS `new_price`,`c`.`hy_name` AS `hy_name`,`a`.`day0` AS `day0`,`a`.`min` AS `min`,`a`.`max` AS `max`,`a`.`percent` AS `percent`,`a`.`fyyx` AS `fyyx`,`a`.`zlzxh` AS `zlzxh`,`a`.`fx` AS `fx`,`a`.`smx` AS `smx`,`a`.`dwzq` AS `dwzq`,`a`.`ksls` AS `ksls`,`a`.`jdz` AS `jdz`,`a`.`doc` AS `doc`,`a`.`count` AS `count`,`a`.`turnover_rate` AS `turnover_rate`,`a`.`created_at` AS `created_at` from ((`analy_tdx_tbl` `a` join `good_shares_tbl` `b` on((`b`.`code` = `a`.`code`))) join `shares_info_tbl` `c` on((`c`.`code` = `a`.`code`))) order by `a`.`day0` desc;

-- ----------------------------
-- View structure for analy_up_view
-- ----------------------------
DROP VIEW IF EXISTS `analy_up_view`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `analy_up_view` AS select `a`.`id` AS `id`,`a`.`code` AS `code`,`a`.`name` AS `name`,`a`.`day0` AS `day0`,`a`.`day_str` AS `day_str`,`a`.`price` AS `price`,`a`.`new_price` AS `new_price`,`a`.`percent` AS `percent`,`a`.`turnover_rate` AS `turnover_rate`,`a`.`doc` AS `doc`,`a`.`created_at` AS `created_at`,`c`.`hy_name` AS `hy_name` from ((`analy_up_tbl` `a` join `good_shares_tbl` `b` on((`b`.`code` = `a`.`code`))) join `shares_info_tbl` `c` on((`c`.`code` = `a`.`code`)));

-- ----------------------------
-- View structure for cd_view
-- ----------------------------
DROP VIEW IF EXISTS `cd_view`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `cd_view` AS select `a`.`id` AS `id`,`a`.`code` AS `code`,`a`.`name` AS `name`,`a`.`day0` AS `day0`,`a`.`day_str` AS `day_str`,`a`.`price` AS `price`,`a`.`new_price` AS `new_price`,`a`.`percent` AS `percent`,`a`.`turnover_rate` AS `turnover_rate`,`a`.`score` AS `score`,`a`.`doc` AS `doc`,`a`.`created_at` AS `created_at`,`c`.`hy_name` AS `hy_name` from (`analy_cd_tbl` `a` join `shares_info_tbl` `c` on((`c`.`code` = `a`.`code`)));

-- ----------------------------
-- View structure for ma20_view
-- ----------------------------
DROP VIEW IF EXISTS `ma20_view`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `ma20_view` AS select distinct `a`.`code` AS `code`,`b`.`name` AS `name`,`b`.`price` AS `price`,`e`.`ma20` AS `ma20`,`b`.`peg` AS `peg`,`b`.`peg_shares` AS `peg_shares`,`b`.`hy_name` AS `hy_name`,`c`.`nick_name` AS `nick_name`,`e`.`day0_str` AS `day0_str` from (((`shares_db`.`shares_watch_tbl` `a` join `shares_db`.`shares_info_tbl` `b` on((`b`.`code` = `a`.`code`))) join `shares_db`.`wx_userinfo` `c` on((convert(`c`.`openid` using utf8mb4) = `a`.`open_id`))) join `shares_db`.`shares_daily_tbl` `e` on(((`e`.`code` = `a`.`code`) and (`b`.`price` >= `e`.`ma20`) and `e`.`day0` in (select `t`.`day0` from (select `shares_db`.`shares_daily_tbl`.`day0` AS `day0` from `shares_db`.`shares_daily_tbl` where (`shares_db`.`shares_daily_tbl`.`code` = 'sh600000') order by `shares_db`.`shares_daily_tbl`.`day0` desc limit 0,1) `t`))));

-- ----------------------------
-- View structure for msg_code_view
-- ----------------------------
DROP VIEW IF EXISTS `msg_code_view`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`%` SQL SECURITY DEFINER VIEW `msg_code_view` AS select `a`.`code` AS `code`,'用户监听' AS `tp` from `shares_watch_tbl` `a` union select `b`.`code` AS `code`,'超跌反弹' AS `tp` from `analy_cd_tbl` `b` union select `c`.`code` AS `code`,'放巨量' AS `tp` from `analy_fl_tbl` `c` union select `d`.`code` AS `code`,'趋势up' AS `tp` from `analy_up_tbl` `d` union select `e`.`code` AS `code`,'龙虎榜' AS `tp` from `analy_lhb_tbl` `e` union select `f`.`code` AS `code`,'底部十字星' AS `tp` from `analy_dbszx_tbl` `f`;

SET FOREIGN_KEY_CHECKS = 1;

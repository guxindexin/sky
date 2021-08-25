BEGIN;
-- 用户
INSERT INTO system_user (username, password, nickname, tel, email, sex, avatar, status, remark, role, dept, id, create_time, update_time, delete_time) VALUES ('lanyulei', '$2a$10$7AkPs/BP4h3JQiK0K2vPLeh/qa.soqsiCllhjwvHYkfEnHo./kQo6', '', '', '', 0, '', false, '', null, null, 1, '2021-08-25 16:33:23.108187 +00:00', '2021-08-25 16:33:23.108187 +00:00', null);

-- 菜单
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (1, '/home', 'home', 'home/index', 'message.router.home', '', false, true, true, false, '{"admin", "test"}', 'iconfont icon-shouye', 0, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (2, '/system', 'system', 'layout/routerView/parent', 'message.router.system', '', false, true, false, false, '{"admin"}', 'iconfont icon-xitongshezhi', 0, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (3, '/system/menu', 'systemMenu', 'system/menu/index', 'message.router.systemMenu', '', false, true, false, false, '{"admin"}', 'iconfont icon-caidan', 2, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (4, '/system/user', 'systemUser', 'system/user/index', 'message.router.systemUser', '', false, true, false, false, '{"admin"}', 'iconfont icon-icon-', 2, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (5, '/limits', 'limits', 'layout/routerView/parent', 'message.router.limits', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-quanxian', 0, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, parent, type, create_time, update_time) VALUES (6, '/limits/backEnd', 'limitsBackEnd', 'layout/routerView/parent', 'message.router.limitsBackEnd', '', false, true, false, false, '{"admin", "test"}', 5, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, parent, type, create_time, update_time) VALUES (7, '/limits/backEnd/page', 'limitsBackEndEndPage', 'limits/backEnd/page/index', 'message.router.limitsBackEndEndPage', '', false, true, false, false, '{"admin", "test"}', 6, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (8, '/menu', 'menu', 'layout/routerView/parent', 'message.router.menu', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-caidan', 0, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (9, '/menu/menu1', 'menu1', 'layout/routerView/parent', 'message.router.menu1', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-caidan', 8, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (10, '/menu/menu1/menu11', 'menu11', 'menu/menu1/menu11/index', 'message.router.menu11', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-caidan', 9, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (11, '/menu/menu1/menu12', 'menu12', 'layout/routerView/parent', 'message.router.menu12', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-caidan', 9, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (12, '/menu/menu1/menu12/menu121', 'menu121', 'menu/menu1/menu12/menu121/index', 'message.router.menu121', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-caidan', 11, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (13, '/menu/menu1/menu12/menu122', 'menu122', 'menu/menu1/menu12/menu122/index', 'message.router.menu122', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-caidan', 11, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (14, '/menu/menu1/menu13', 'menu13', 'menu/menu1/menu13/index', 'message.router.menu13', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-caidan', 9, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (15, '/menu/menu2', 'menu2', 'menu/menu2/index', 'message.router.menu2', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-caidan', 8, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (16, '/fun', 'funIndex', 'layout/routerView/parent', 'message.router.funIndex', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-crew_feature', 0, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (17, '/fun/tagsView', 'funTagsView', 'fun/tagsView/index', 'message.router.funTagsView', '', false, true, false, false, '{"admin", "test"}', 'el-icon-thumb', 16, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (18, '/chart', 'chartIndex', 'chart/index', 'message.router.chartIndex', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-ico_shuju', 0, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (19, '/personal', 'personal', 'personal/index', 'message.router.personal', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-gerenzhongxin', 0, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (20, '/tools', 'tools', 'tools/index', 'message.router.tools', '', false, true, false, false, '{"admin", "test"}', 'iconfont icon-gongju', 0, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (21, '/link', 'layoutLinkView', 'layout/routerView/link', 'message.router.layoutLinkView', 'https://element-plus.gitee.io/#/zh-CN/component/installation', false, false, false, false, '{"admin"}', 'iconfont icon-caozuo-wailian', 0, 1, now(), now());
INSERT INTO system_menu (id, path, name, component, title, is_link, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, create_time, update_time) VALUES (22, '/iframes', 'layoutIfameView', 'layout/routerView/iframe', 'message.router.layoutIfameView', 'https://gitee.com/lyt-top/vue-next-admin', false, true, true, true, '{"admin"}', 'iconfont icon-neiqianshujuchucun', 0, 1, now(), now());
COMMIT;
-- 执行结束

BEGIN;
-- 用户
INSERT INTO `system_user` (username, password, nickname, tel, email, sex, avatar, status, remark, role, dept, id, create_time, update_time, delete_time) VALUES ('lanyulei', '$2a$10$7AkPs/BP4h3JQiK0K2vPLeh/qa.soqsiCllhjwvHYkfEnHo./kQo6', '兰玉磊', '', '', 0, '', false, '', null, null, 1, '2021-08-25 16:33:23.108187 +00:00', '2021-08-25 16:33:23.108187 +00:00', null);

-- 菜单
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, id, create_time, update_time, delete_time, sort) VALUES ('/home', 'home', 'home/index', null, '首页', '', false, true, true, false, '{admin,test}', 'iconfont icon-shouye', 0, 2, 1, '2021-08-25 16:18:47.782596 +00:00', '2021-08-25 16:18:47.782596 +00:00', null, 1);
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, id, create_time, update_time, delete_time, sort) VALUES ('/system', 'system', 'layout/routerView/parent', '/system/menu', '系统设置', '', false, true, false, false, '{admin,test}', 'iconfont icon-xitongshezhi', 0, 1, 2, '2021-08-25 16:18:47.782596 +00:00', '2021-08-25 16:18:47.782596 +00:00', null, 1);
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, id, create_time, update_time, delete_time, sort) VALUES ('/system/menu', 'systemMenu', 'system/menu/index', null, '菜单管理', '', false, true, false, false, '{admin,test}', 'iconfont icon-caidan', 2, 2, 3, '2021-08-25 16:18:47.782596 +00:00', '2021-08-25 16:18:47.782596 +00:00', null, 1);
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, id, create_time, update_time, delete_time, sort) VALUES ('/system/user', 'systemUser', 'system/user/index', null, '用户管理', '', false, true, false, false, '{admin,test}', 'iconfont icon-icon-', 2, 2, 4, '2021-08-25 16:18:47.782596 +00:00', '2021-08-25 16:18:47.782596 +00:00', null, 1);

COMMIT;
-- 执行结束

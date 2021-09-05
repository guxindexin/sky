BEGIN;
-- 用户
INSERT INTO system_user (username, password, nickname, tel, email, sex, status, is_admin, remark, role, create_time, update_time, delete_time) VALUES ('lanyulei', '$2a$10$7AkPs/BP4h3JQiK0K2vPLeh/qa.soqsiCllhjwvHYkfEnHo./kQo6', '兰玉磊', '', '', 0, true, true, '', null, '2021-08-25 16:33:23.108187 +00:00', '2021-08-25 16:33:23.108187 +00:00', null);

-- 菜单
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort, create_time, update_time, delete_time) VALUES ('/home', 'Home', 'home/index', '', '首页', '', false, true, true, false, '{admin,test}', 'iconfont icon-shouye', 0, 2, 0, '2021-09-02 15:30:08.646648 +00:00', '2021-09-02 15:30:08.646648 +00:00', null);
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort, create_time, update_time, delete_time) VALUES ('/system', 'System', 'layout/routerView/parent', '', '系统设置', '', false, true, false, false, '{admin,test}', 'iconfont icon-xitongshezhi', 0, 1, 2, '2021-09-02 15:30:15.576879 +00:00', '2021-09-02 15:30:15.576879 +00:00', null);
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort, create_time, update_time, delete_time) VALUES ('/system/menu', 'SystemMenu', 'system/menu/index', '', '菜单管理', '', false, true, false, false, '{admin,test}', 'el-icon-tickets', 2, 2, 15, '0001-01-01 00:00:00.000000 +00:00', '2021-09-02 15:54:09.531571 +00:00', null);
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort, create_time, update_time, delete_time) VALUES ('/system/user', 'SystemUser', 'system/user/index', '', '用户管理', '', false, true, false, false, '{test,admin}', 'el-icon-user', 2, 2, 1, '0001-01-01 00:00:00.000000 +00:00', '2021-09-02 15:54:16.157875 +00:00', null);
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort, create_time, update_time, delete_time) VALUES ('/system/role', 'SystemRole', 'system/role/index', '', '角色管理', '', false, true, false, false, '{admin,test}', 'el-icon-picture-outline-round', 2, 2, 5, '0001-01-01 00:00:00.000000 +00:00', '2021-09-02 15:54:14.040620 +00:00', null);
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort, create_time, update_time, delete_time) VALUES ('/system/api', 'SystemAPI', 'system/api/index', '', '接口管理', '', false, true, false, false, '{test,admin}', 'el-icon-s-promotion', 2, 2, 10, '0001-01-01 00:00:00.000000 +00:00', '2021-09-02 15:59:37.740717 +00:00', null);
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort, create_time, update_time, delete_time) VALUES ('/personal', 'Personal', 'personal/index', '', '个人中心', '', true, true, false, false, '{admin,test}', 'el-icon-help', 2, 2, 99, '2021-09-02 16:30:39.053882 +00:00', '2021-09-02 16:30:39.053882 +00:00', null);
INSERT INTO system_menu (path, name, component, redirect, title, hyperlink, is_hide, is_keep_alive, is_affix, is_iframe, auth, icon, parent, type, sort, create_time, update_time, delete_time) VALUES ('', '', '', '', '新建用户', '', false, false, false, false, '{system:user:create}', '', 4, 3, 0, '2021-09-02 17:04:58.506348 +00:00', '2021-09-02 17:04:58.506348 +00:00', null);

-- API分组
INSERT INTO system_api_group (app, remark, create_time, update_time) VALUES ('系统管理', '', '2021-09-05 10:25:43.667889 +00:00', '2021-09-05 10:49:02.592078 +00:00');

-- API接口
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('菜单对应的按钮', '/api/v1/system/menu/button/:id', 'GET', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('新建用户', '/api/v1/system/user', 'POST', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('角色列表', '/api/v1/system/role', 'GET', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('保存角色', '/api/v1/system/role', 'POST', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('删除菜单', '/api/v1/system/menu/:id', 'DELETE', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('删除角色', '/api/v1/system/role/:id', 'DELETE', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('删除用户', '/api/v1/system/user/:id', 'DELETE', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('接口列表', '/api/v1/system/api', 'GET', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('编辑用户', '/api/v1/system/user/:id', 'PUT', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('删除接口', '/api/v1/system/api/:id', 'DELETE', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('批量删除菜单', '/api/v1/system/menu/batch', 'DELETE', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('用户列表', '/api/v1/system/user', 'GET', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('保存菜单', '/api/v1/system/menu', 'POST', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('保存接口', '/api/v1/system/api', 'POST', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('菜单树结构', '/api/v1/system/menu/tree', 'GET', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');
INSERT INTO system_api (title, url, method, "group", remark, create_time, update_time) VALUES ('用户详情', '/api/v1/system/user/info', 'GET', 1, '', '2021-09-04 21:44:18.508000 +00:00', '2021-09-04 21:44:18.508000 +00:00');

COMMIT;
-- 执行结束

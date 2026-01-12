INSERT INTO registered_channel VALUES ('2570cc0e-deaf-401c-b9e6-eaaae8173792','neopay','1','2024-10-14 09:10:12','2024-10-15 08:32:08');
INSERT INTO registered_channel VALUES ('399e7c63-674c-40af-bf88-48a246a2f228','fdstest','1','2024-12-23 07:23:09','2024-12-23 07:23:09');
INSERT INTO registered_channel VALUES ('81c1bc0c-4d60-4e45-9056-86d1772ba719','neo','1','2024-10-14 09:10:12','2024-10-15 08:32:08');
INSERT INTO registered_channel VALUES ('a4d92091-dd37-41c7-b19a-09cada495035','yoopay','1','2024-10-14 09:10:12','2024-10-15 08:32:08');
INSERT INTO registered_channel VALUES ('e8d07cd3-1380-4c87-8996-5f5c0485c304','tsel','1','2025-02-04 04:13:50','2025-02-04 04:13:50');

INSERT INTO roles VALUES (1, 'Admin', '2021-09-28 21:01:23', '2024-01-24 01:07:06', 'Admin can access all menu');
INSERT INTO roles VALUES (2, 'User Maker', '2021-09-28 21:01:23', '2024-01-24 01:07:07', 'User - Maker can only access to menu: Blacklist');
INSERT INTO roles VALUES (3, 'User Checker', '2021-09-28 21:01:25', '2024-02-13 08:16:14', 'User - Checker can only access to menu: Blacklist');

INSERT INTO flag VALUES ('aedcac7f-064b-43d7-b866-47a2875eec63', 'Warning', '2021-09-29 14:50:28', '2021-09-29 14:50:28');
INSERT INTO flag VALUES ('b53ce142-b2fd-4634-a964-801dd7bfaf3b', 'Rejected', '2021-09-21 22:01:55', '2021-09-21 22:01:55');
INSERT INTO flag VALUES ('e2077d0f-aaf5-4520-95d7-e4b097d2f3a5', 'Allowed', '2021-09-29 14:50:28', '2021-09-29 14:50:28');

INSERT INTO users VALUES ('28f984b7-a9c6-49be-8004-cfa496da1c13', 1, 'maker.account@gmail.com', '$2a$10$UVcouDROHNDBNgBm0n4djOGrakzCrsH8lvlsa3RJ5lGm/zCR1yXpW', 'Account for checker role', 'https://icons.veryicon.com/png/o/miscellaneous/standard/avatar-15.png', 'M', 0, '2025-03-26 14:41:45', '', '2025-03-26 14:41:45', NULL, NULL, NULL);
INSERT INTO users VALUES ('21a18579-a12a-4f99-a887-0e8167f5adb5', 1, 'checker.account@gmail.com', '$2a$10$vFtUlOJtJu8oUaAIAp/3aeWqFn2XiLAm/s1XyfWQaeuAaxolRSqDm', 'Account for maker role', 'https://icons.veryicon.com/png/o/miscellaneous/standard/avatar-15.png', 'M', 0, '2025-03-26 14:42:02', '', '2025-03-26 14:42:02', NULL, NULL, NULL);

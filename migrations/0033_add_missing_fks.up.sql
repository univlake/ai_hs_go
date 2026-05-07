-- 集中补齐 0001~0028 遗漏的外键，对齐 0032 special_equipment.special_type_id 风格。
--
-- ON DELETE 策略：
--   companies(id)         RESTRICT   租户根，禁止级联消除业务数据
--   字典表 (lookup)        SET NULL   字典消失不应清空业务行
--   父子明细 (kids/scopes) CASCADE    主表删除时子表跟随
--
-- 显式跳过（设计上有意为之，源迁移注释里说明）：
--   * 0011 regulation_library_types.company_id           字典全局共用，目前都填 0
--   * 0017 design_units / evaluation_units / skill_units 不挂字典 FK（沿用 Laravel）
--   * 0020 evaluation_unit_id / design_unit_id / skill_unit* 软删除时不级联噪音
--   * 0025 *_templates.pid 自引用                         pid=0 哨兵表示顶级
--
-- DEFAULT 0 哨兵清理：
--   risk_logs / regulation_libraries / inherent_risk_logs / float_risk_logs
--   的 company_id 移除 DEFAULT 0。预检确认本地无 company_id=0 行；其它环境
--   先 cleanup 再 migrate。

-- 1. companies 字典外键 ─────────────────────────────────────────
ALTER TABLE companies
    ADD CONSTRAINT companies_enterprise_type_id_fkey
        FOREIGN KEY (enterprise_type_id) REFERENCES enterprise_types(id) ON DELETE SET NULL;

-- operational_type_id 是 NOT NULL DEFAULT 1，不能 SET NULL；用 RESTRICT 防误删字典。
ALTER TABLE companies
    ADD CONSTRAINT companies_operational_type_id_fkey
        FOREIGN KEY (operational_type_id) REFERENCES operational_types(id) ON DELETE RESTRICT;

ALTER TABLE companies
    ADD CONSTRAINT companies_area_id_fkey
        FOREIGN KEY (area_id) REFERENCES base_areas(area_id) ON DELETE SET NULL;

-- 2. 父-子明细（CASCADE）─────────────────────────────────────────
ALTER TABLE skill_unit_kids
    ADD CONSTRAINT skill_unit_kids_skill_unit_id_fkey
        FOREIGN KEY (skill_unit_id) REFERENCES skill_units(id) ON DELETE CASCADE;

ALTER TABLE internal_meeting_kids
    ADD CONSTRAINT internal_meeting_kids_internal_meeting_id_fkey
        FOREIGN KEY (internal_meeting_id) REFERENCES internal_meetings(id) ON DELETE CASCADE;

ALTER TABLE training_file_logs
    ADD CONSTRAINT training_file_logs_training_file_id_fkey
        FOREIGN KEY (training_file_id) REFERENCES training_files(id) ON DELETE CASCADE;

ALTER TABLE training_video_logs
    ADD CONSTRAINT training_video_logs_training_video_id_fkey
        FOREIGN KEY (training_video_id) REFERENCES training_videos(id) ON DELETE CASCADE;

-- 3. 通知子表 company_group_id（lookup → SET NULL）──────────────
ALTER TABLE office_notify_scopes
    ADD CONSTRAINT office_notify_scopes_company_group_id_fkey
        FOREIGN KEY (company_group_id) REFERENCES company_groups(id) ON DELETE SET NULL;

ALTER TABLE new_notify_scopes
    ADD CONSTRAINT new_notify_scopes_company_group_id_fkey
        FOREIGN KEY (company_group_id) REFERENCES company_groups(id) ON DELETE SET NULL;

ALTER TABLE meeting_notices
    ADD CONSTRAINT meeting_notices_company_group_id_fkey
        FOREIGN KEY (company_group_id) REFERENCES company_groups(id) ON DELETE SET NULL;

-- 4. 信息稿审核人（user lookup）────────────────────────────────
-- 历史脏数据（早期冒烟测试遗留 next_reviewer 指向不存在的 user）会阻塞 FK
-- 创建。NOT VALID 让新插入受约束，老数据不动；运维确认无问题后执行：
--   ALTER TABLE information_drafts VALIDATE CONSTRAINT information_drafts_next_reviewer_fkey;
ALTER TABLE information_drafts
    ADD CONSTRAINT information_drafts_next_reviewer_fkey
        FOREIGN KEY (next_reviewer) REFERENCES users(id) ON DELETE SET NULL NOT VALID;

-- 5. 法律法规库 → 类型（lookup → SET NULL）────────────────────
ALTER TABLE regulation_libraries
    ADD CONSTRAINT regulation_libraries_regulation_library_type_id_fkey
        FOREIGN KEY (regulation_library_type_id) REFERENCES regulation_library_types(id) ON DELETE SET NULL;

-- 6. 集团关注关系 ─────────────────────────────────────────────
-- 主体公司不能直接删（有业务数据），关注的目标删了关系自然清掉。
ALTER TABLE focus_companies
    ADD CONSTRAINT focus_companies_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;
ALTER TABLE focus_companies
    ADD CONSTRAINT focus_companies_focus_id_fkey
        FOREIGN KEY (focus_id) REFERENCES companies(id) ON DELETE CASCADE;

-- 7. 考核五张表：company + 模板 ───────────────────────────────
ALTER TABLE daily_checks
    ADD CONSTRAINT daily_checks_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT,
    ADD CONSTRAINT daily_checks_template_id_fkey
        FOREIGN KEY (daily_check_template_id) REFERENCES daily_check_templates(id) ON DELETE RESTRICT;

ALTER TABLE accident_checks
    ADD CONSTRAINT accident_checks_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT,
    ADD CONSTRAINT accident_checks_template_id_fkey
        FOREIGN KEY (accident_check_template_id) REFERENCES accident_check_templates(id) ON DELETE RESTRICT;

ALTER TABLE add_points
    ADD CONSTRAINT add_points_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT,
    ADD CONSTRAINT add_points_template_id_fkey
        FOREIGN KEY (add_points_template_id) REFERENCES add_points_templates(id) ON DELETE RESTRICT;

ALTER TABLE system_checks
    ADD CONSTRAINT system_checks_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE norm_checks
    ADD CONSTRAINT norm_checks_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

-- 8. 培训流水 company_id ──────────────────────────────────────
ALTER TABLE training_file_logs
    ADD CONSTRAINT training_file_logs_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE training_video_logs
    ADD CONSTRAINT training_video_logs_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

-- 9. DEFAULT 0 哨兵清理 + FK ──────────────────────────────────
ALTER TABLE risk_logs
    ALTER COLUMN company_id DROP DEFAULT,
    ADD CONSTRAINT risk_logs_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE regulation_libraries
    ALTER COLUMN company_id DROP DEFAULT,
    ADD CONSTRAINT regulation_libraries_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE inherent_risk_logs
    ALTER COLUMN company_id DROP DEFAULT,
    ADD CONSTRAINT inherent_risk_logs_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

ALTER TABLE float_risk_logs
    ALTER COLUMN company_id DROP DEFAULT,
    ADD CONSTRAINT float_risk_logs_company_id_fkey
        FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE RESTRICT;

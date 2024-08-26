-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE o_device
(
    id              VARCHAR(32) NOT NULL,
    created_by      VARCHAR(32),
    created_time    TIMESTAMP,
    updated_by      VARCHAR(32),
    updated_time    TIMESTAMP,
    name            VARCHAR(255),
    type            INTEGER,
    status          INTEGER,
    parent_id       VARCHAR(32),
    group_id        VARCHAR(32),
    product_id      VARCHAR(32),
    protocol_config TEXT,
    identifier      VARCHAR(255),
    enabled         INTEGER,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_device IS '设备';
COMMENT ON COLUMN o_device.id IS '唯一标识(md5计算)';
COMMENT ON COLUMN o_device.created_by IS '创建人';
COMMENT ON COLUMN o_device.created_time IS '创建时间';
COMMENT ON COLUMN o_device.updated_by IS '更新人';
COMMENT ON COLUMN o_device.updated_time IS '更新时间';
COMMENT ON COLUMN o_device.name IS '名称';
COMMENT ON COLUMN o_device.type IS '类型(1设备，2网关)';
COMMENT ON COLUMN o_device.status IS '状态(0:离线:1在线,2:告警)';
COMMENT ON COLUMN o_device.parent_id IS '父id';
COMMENT ON COLUMN o_device.group_id IS '分组id';
COMMENT ON COLUMN o_device.product_id IS '产品id';
COMMENT ON COLUMN o_device.protocol_config IS '协议配置(json)';
COMMENT ON COLUMN o_device.identifier IS '标识';
COMMENT ON COLUMN o_device.enabled IS '启用禁用';



CREATE TABLE f_point_data
(
    id    VARCHAR(32) NOT NULL,
    ts    TIMESTAMP   NOT NULL,
    key   VARCHAR(32) NOT NULL,
    value FLOAT8,
    PRIMARY KEY (id, ts, key)
);

COMMENT ON TABLE f_point_data IS '点位数据';
COMMENT ON COLUMN f_point_data.id IS '点位id';
COMMENT ON COLUMN f_point_data.ts IS '创建时间';
COMMENT ON COLUMN f_point_data.value IS '值';



CREATE TABLE f_device_data
(
    id                  VARCHAR(32)  NOT NULL,
    device_identifier   VARCHAR(128) NOT NULL,
    property_identifier VARCHAR(128) NOT NULL,
    ts                  TIMESTAMP    NOT NULL,
    vtype               int,
    unit                varchar(32),
    f_value             FLOAT8,
    i_value             integer,
    t_value             TIMESTAMP,
    s_value             TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE f_device_data IS '设备数据';
COMMENT ON COLUMN f_device_data.id IS '唯一id(device_identifier+property_identifier+ts md5)';
COMMENT ON COLUMN f_device_data.ts IS '创建时间';
COMMENT ON COLUMN f_device_data.vtype IS '值类型1 float,2:int,3:ts,4:string';
COMMENT ON COLUMN f_device_data.unit IS '单位';
COMMENT ON COLUMN f_device_data.f_value IS 'float值';
COMMENT ON COLUMN f_device_data.i_value IS 'int值';
COMMENT ON COLUMN f_device_data.t_value IS '时间值';
COMMENT ON COLUMN f_device_data.s_value IS 'string值';



CREATE TABLE o_point
(
    id         VARCHAR(32)  NOT NULL,
    gateway_id VARCHAR(255) NOT NULL,
    device_id  VARCHAR(255) NOT NULL,
    name       VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_point IS '点位';
COMMENT ON COLUMN o_point.id IS '唯一标识';
COMMENT ON COLUMN o_point.gateway_id IS '网关id(name->md5)';
COMMENT ON COLUMN o_point.device_id IS '设备id(name->md5)';
COMMENT ON COLUMN o_point.name IS '名称';

alter table o_point
    add column status integer not null default 0;



CREATE TABLE o_tag
(
    id       VARCHAR(32)  NOT NULL,
    rel_id   VARCHAR(255) NOT NULL,
    key      VARCHAR(255) NOT NULL,
    value    VARCHAR(255) NOT NULL,
    editable INTEGER      NOT NULL,
    rel_type INTEGER,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_tag IS 'tag实例';
COMMENT ON COLUMN o_tag.id IS '唯一标识';
COMMENT ON COLUMN o_tag.rel_id IS '关联id';
COMMENT ON COLUMN o_tag.key IS '名称';
COMMENT ON COLUMN o_tag.value IS '值';
COMMENT ON COLUMN o_tag.editable IS '是否可编辑（导入的不可编辑）';
COMMENT ON COLUMN o_tag.rel_type IS '(1:网关、2:设备、3:点位';

alter table o_tag
    add user_id varchar(32) NOT NULL default '';
comment on column o_tag.user_id is '用户id';



CREATE TABLE o_product
(
    id           VARCHAR(32)  NOT NULL,
    created_by   VARCHAR(32)  NOT NULL,
    created_time TIMESTAMP    NOT NULL,
    updated_by   VARCHAR(32)  NOT NULL,
    updated_time TIMESTAMP    NOT NULL,
    name         VARCHAR(255) NOT NULL,
    json_data    TEXT         NOT NULL,
    vendor       VARCHAR(255) NOT NULL,
    model        VARCHAR(255) NOT NULL,
    version      VARCHAR(255) NOT NULL,
    identifier   VARCHAR(32),
    categories   VARCHAR(255),
    descript     TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_product IS '产品表';
COMMENT ON COLUMN o_product.id IS '唯一标识(nanoid)';
COMMENT ON COLUMN o_product.created_by IS '创建人';
COMMENT ON COLUMN o_product.created_time IS '创建时间';
COMMENT ON COLUMN o_product.updated_by IS '更新人';
COMMENT ON COLUMN o_product.updated_time IS '更新时间';
COMMENT ON COLUMN o_product.name IS '产品名称';
COMMENT ON COLUMN o_product.json_data IS '产品物模型';
COMMENT ON COLUMN o_product.vendor IS '厂商';
COMMENT ON COLUMN o_product.model IS '型号';
COMMENT ON COLUMN o_product.version IS '版本';
COMMENT ON COLUMN o_product.identifier IS '标识';
COMMENT ON COLUMN o_product.categories IS '分类（A/B/C）';
COMMENT ON COLUMN o_product.descript IS '描述';

alter table o_product
    add column type integer not null default 1;
comment on column o_product.type is '类型(1:设备:2:网关)';



CREATE TABLE o_device_model
(
    id              VARCHAR(32)  NOT NULL,
    created_by      VARCHAR(32)  NOT NULL,
    created_time    TIMESTAMP    NOT NULL,
    updated_by      VARCHAR(32)  NOT NULL,
    updated_time    TIMESTAMP    NOT NULL,
    name            VARCHAR(255) NOT NULL,
    descript        VARCHAR(255) NOT NULL,
    categories      VARCHAR(255) NOT NULL,
    json_data       TEXT         NOT NULL,
    service_script  TEXT         NOT NULL,
    property_script TEXT         NOT NULL,
    event_script    TEXT         NOT NULL,
    cover_file      TEXT         NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_device_model IS '物模型';
COMMENT ON COLUMN o_device_model.id IS '唯一标识';
COMMENT ON COLUMN o_device_model.created_by IS '创建人';
COMMENT ON COLUMN o_device_model.created_time IS '创建时间';
COMMENT ON COLUMN o_device_model.updated_by IS '更新人';
COMMENT ON COLUMN o_device_model.updated_time IS '更新时间';
COMMENT ON COLUMN o_device_model.name IS '名称';
COMMENT ON COLUMN o_device_model.descript IS '描述';
COMMENT ON COLUMN o_device_model.categories IS '分类（A/B/C 手动填写)';
COMMENT ON COLUMN o_device_model.json_data IS '物模型数据(json定义,属性、服务、事件)';
COMMENT ON COLUMN o_device_model.service_script IS '服务数据转换脚本处理';
COMMENT ON COLUMN o_device_model.property_script IS '属性数据转换脚本处理';
COMMENT ON COLUMN o_device_model.event_script IS '事件转换脚本处理';
COMMENT ON COLUMN o_device_model.cover_file IS '封面base64';



CREATE TABLE o_model_meta
(
    id         VARCHAR(32) NOT NULL,
    name       VARCHAR(255),
    identifier VARCHAR(255),
    json_data  TEXT,
    meta_type  VARCHAR(20),
    category   VARCHAR(255),
    type       VARCHAR(255),
    tag        int8        not null default 0,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_model_meta IS '物模型元数据表';
COMMENT ON COLUMN o_model_meta.id IS '唯一标识';
COMMENT ON COLUMN o_model_meta.name IS '名称';
COMMENT ON COLUMN o_model_meta.identifier IS '标识符';
COMMENT ON COLUMN o_model_meta.json_data IS '元数据';
COMMENT ON COLUMN o_model_meta.meta_type IS '类型(attribute,service,event)';
COMMENT ON COLUMN o_model_meta.type IS '品类';
COMMENT ON COLUMN o_model_meta.tag IS '更新tag';


CREATE TABLE o_alarm_rule
(
    id        VARCHAR(255) NOT NULL,
    create_at TIMESTAMP    NOT NULL,
    update_at TIMESTAMP    NOT NULL,
    create_id VARCHAR(255) NOT NULL,
    update_id VARCHAR(255) NOT NULL,
    name      VARCHAR(255) NOT NULL,
    status    INTEGER      NOT NULL,
    content   text         NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_alarm_rule IS '告警规则表';
COMMENT ON COLUMN o_alarm_rule.id IS '唯一标识';
COMMENT ON COLUMN o_alarm_rule.create_at IS '创建时间';
COMMENT ON COLUMN o_alarm_rule.update_at IS '修改时间';
COMMENT ON COLUMN o_alarm_rule.create_id IS '创建人id';
COMMENT ON COLUMN o_alarm_rule.update_id IS '修改人id';
COMMENT ON COLUMN o_alarm_rule.name IS '名称';
COMMENT ON COLUMN o_alarm_rule.status IS '状态 0:未启用1:启用';
COMMENT ON COLUMN o_alarm_rule.content IS '内容';

CREATE TABLE o_device_mirror
(
    id           VARCHAR(32) NOT NULL,
    json_data    TEXT,
    updated_time TIMESTAMP,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_device_mirror IS '设备孪生';
COMMENT ON COLUMN o_device_mirror.id IS '唯一标识';
COMMENT ON COLUMN o_device_mirror.json_data IS '当前数据json';
COMMENT ON COLUMN o_device_mirror.updated_time IS '更新时间';



CREATE TABLE o_sim_device
(
    id              VARCHAR(32) NOT NULL,
    created_by      VARCHAR(32),
    created_time    TIMESTAMP,
    updated_by      VARCHAR(32),
    updated_time    TIMESTAMP,
    name            VARCHAR(255),
    type            INTEGER,
    parent_id       VARCHAR(32),
    group_id        VARCHAR(32),
    product_id      VARCHAR(32),
    protocol_config TEXT,
    identifier      VARCHAR(255),
    enabled         INTEGER,
    rule_data       TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_sim_device IS '模拟设备';
COMMENT ON COLUMN o_sim_device.id IS '唯一标识(md5计算)';
COMMENT ON COLUMN o_sim_device.created_by IS '创建人';
COMMENT ON COLUMN o_sim_device.created_time IS '创建时间';
COMMENT ON COLUMN o_sim_device.updated_by IS '更新人';
COMMENT ON COLUMN o_sim_device.updated_time IS '更新时间';
COMMENT ON COLUMN o_sim_device.name IS '名称';
COMMENT ON COLUMN o_sim_device.type IS '类型(1设备，2网关)';
COMMENT ON COLUMN o_sim_device.parent_id IS '父id';
COMMENT ON COLUMN o_sim_device.group_id IS '分组id';
COMMENT ON COLUMN o_sim_device.product_id IS '产品id';
COMMENT ON COLUMN o_sim_device.protocol_config IS '协议配置(json)';
COMMENT ON COLUMN o_sim_device.identifier IS '标识';
COMMENT ON COLUMN o_sim_device.enabled IS '启用禁用';
COMMENT ON COLUMN o_sim_device.rule_data IS '规则数据JSON';



CREATE TABLE f_sim_device_data
(
    id                  VARCHAR(32) NOT NULL,
    device_identifier   VARCHAR(32) NOT NULL,
    property_identifier VARCHAR(32) NOT NULL,
    ts                  TIMESTAMP   NOT NULL,
    f_value             FLOAT8,
    i_value             integer,
    t_value             TIMESTAMP,
    s_value             TEXT,
    PRIMARY KEY (id, ts)
);

COMMENT ON TABLE f_sim_device_data IS '设备数据';
COMMENT ON COLUMN f_sim_device_data.id IS '唯一id(device_identifier+property_identifier+ts md5)';
COMMENT ON COLUMN f_sim_device_data.ts IS '创建时间';
COMMENT ON COLUMN f_sim_device_data.f_value IS 'float值';
COMMENT ON COLUMN f_sim_device_data.i_value IS 'int值';
COMMENT ON COLUMN f_sim_device_data.t_value IS '时间值';
COMMENT ON COLUMN f_sim_device_data.s_value IS 'string值';

SELECT create_hypertable('f_sim_device_data', 'ts');

CREATE TABLE o_access_protocol
(
    id           VARCHAR(32) NOT NULL,
    created_by   VARCHAR(32),
    created_time TIMESTAMP,
    updated_by   VARCHAR(32),
    updated_time TIMESTAMP,
    identifier   VARCHAR(255),
    status       INTEGER,
    type         VARCHAR(255),
    properties   TEXT,
    enabled      INTEGER,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_access_protocol IS '数据接入点表';
COMMENT ON COLUMN o_access_protocol.id IS '唯一标识';
COMMENT ON COLUMN o_access_protocol.created_by IS '创建人';
COMMENT ON COLUMN o_access_protocol.created_time IS '创建时间';
COMMENT ON COLUMN o_access_protocol.updated_by IS '更新人';
COMMENT ON COLUMN o_access_protocol.updated_time IS '更新时间';
COMMENT ON COLUMN o_access_protocol.identifier IS '标识名';
COMMENT ON COLUMN o_access_protocol.status IS '状态(0断线，1在线)';
COMMENT ON COLUMN o_access_protocol.type IS '连接方式';
COMMENT ON COLUMN o_access_protocol.properties IS '属性(不同连接方式有不同的属性,json格式）';
COMMENT ON COLUMN o_access_protocol.enabled IS '是否启用(0禁用，1启用';


CREATE TABLE o_user_device
(
    id           VARCHAR(32) NOT NULL,
    user_id      VARCHAR(32),
    device_id    VARCHAR(255),
    created_time TIMESTAMP,
    updated_time TIMESTAMP,
    index        integer,
    locked       integer,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_user_device IS '用户设备管理';
COMMENT ON COLUMN o_user_device.id IS '唯一标识';
COMMENT ON COLUMN o_user_device.user_id IS '用户id';
COMMENT ON COLUMN o_user_device.device_id IS '设备id';
COMMENT ON COLUMN o_user_device.created_time IS '创建时间';
COMMENT ON COLUMN o_user_device.updated_time IS '更新时间';
COMMENT ON COLUMN o_user_device.index IS '索引排序';
COMMENT ON COLUMN o_user_device.locked IS '是否锁定(0,1)';


create table if not exists o_holiday_json
(
    id        varchar(32) not null,
    json_data text,
    primary key (id)
);
comment on table o_holiday_json is '节假日json';
comment on column o_holiday_json.id is 'id(年份，例如2023）';
comment on column o_holiday_json.json_data is 'json数据';


CREATE TABLE o_edge_node
(
    id           VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP,
    PRIMARY KEY (id)
);


create or replace view v_device_tree as
select d.*,
       case d.type when 1 then 0 when 2 then 1 end as           hasChild,
       (select array_to_json(array_agg(e))
        from (select * from o_device where parent_id = d.id) e) children
from o_device d
where parent_id = '0';


create or replace view v_device_identifier_product_json as
select d.identifier, d.status, p.json_data
from o_device d,
     o_product p
where d.product_id = p.id;

create or replace view v_device_with_tag as
select t.*, p.name as product_name
from (SELECT o_device.*,
             array_agg(o_tag.key || ':' || o_tag.value) AS tags
      FROM o_device
               LEFT JOIN o_tag ON o_device.id = o_tag.rel_id
      GROUP BY o_device.id) t,
     o_product p
where t.product_id = p.id;

create or replace view v_device_with_tag_filter as
select t.*, p.name as product_name
from (SELECT o_device.*,
             array_agg(o_tag.key || ':' || o_tag.value)           AS tags,
             string_agg(o_tag.value, ' ') || ' ' || o_device.name AS filter_text
      FROM o_device
               LEFT JOIN o_tag ON o_device.id = o_tag.rel_id
      GROUP BY o_device.id) t,
     o_product p
where t.product_id = p.id;

create or replace view v_point_with_tag as
SELECT o_point.*,
       array_agg(o_tag.key || ':' || o_tag.value) AS tags
FROM o_point
         LEFT JOIN o_tag ON o_point.id = o_tag.rel_id
GROUP BY o_point.id;


CREATE or replace VIEW v_device_point_data AS
SELECT o_device.id   AS id,
       o_device.name AS name,
       o_point.id    AS point_id,
       o_point.name  AS point_name,
       f_point_data.ts,
       f_point_data.key,
       f_point_data.value
FROM o_device
         INNER JOIN o_point ON o_device.id = o_point.device_id
         INNER JOIN f_point_data ON o_point.id = f_point_data.id;


create or replace view v_device_current_data as
SELECT d.id,
       d.name,
       d.tags,
       array_to_json(array_agg((json_build_object('id', f.id, 'name', f.key, 'ts', f.ts, 'value', f.value, 'tags',
                                                  p.tags)))) AS points
FROM v_device_with_tag d
         JOIN v_point_with_tag p ON p.device_id = d.id
         LEFT JOIN (SELECT DISTINCT ON (id, key) id, key, value, ts
                    FROM f_point_data
                    ORDER BY id, key, ts DESC) f ON p.name = f.key
GROUP BY d.id, d.name, d.tags;

create or replace view v_device_info as
select d.*, p.name as product_name
from o_device d,
     o_product p
where d.product_id = p.id;


create or replace view v_tag_with_product_id as
select t.*, o.product_id
from o_tag t
         left join o_device o on t.rel_id = o.id;


create or replace view v_device_with_points as
SELECT d.id,
       d.identifier,
       d.product_id,
       d.enabled,
       d.type,
       d.tags,
       array_to_json(array_agg((json_build_object('id', p.id, 'identifier', p.name, 'tags', p.tags)))) AS points
FROM v_device_with_tag d
         LEFT JOIN v_point_with_tag p ON p.device_id = d.id
GROUP BY d.id, d.identifier, d.product_id, d.enabled, d.type, d.tags;

create or replace view v_user_device_info as
select o1.*, o2.name, o2.type, o2.enabled, o2.identifier, o2.status, o2.tags
from o_user_device o1,
     v_device_with_tag o2,
     o_product p
where o1.device_id = o2.id
  and o2.product_id = p.id;


create or replace view v_nbi_device as
select f.id,
       pp.name as product_name,
       o.name,
       o.identifier,
       o.status,
       p.name  as property_name,
       f.property_identifier,
       f.unit,
       f.ts,
       case
           when f.vtype = 1 then f.f_value
           when f.vtype = 2 then f.i_value
           end as value
from f_device_data f,
     o_device o,
     o_point p,
     o_product pp
where f.id = p.id
  and f.device_identifier = o.identifier
  and o.product_id = pp.id;

create or replace view v_point_info as
select o.id,
       o.gateway_id,
       (select identifier from o_device where id = o.gateway_id) as gateway_identifier,
       o.device_id,
       (select identifier from o_device where id = o.device_id)  as device_identifier,
       o.name,
       f.ts,
       case
           when f.vtype = 1 then f.f_value
           when f.vtype = 2 then f.i_value
           end                                                   as value,
       o.tags
from (select o_point.*, array_agg(o_tag.key || ':' || o_tag.value) AS tags
      from o_point
               left join o_tag on o_point.id = o_tag.rel_id
      group by o_point.id) o
         left join f_device_data f on f.id = o.id;


CREATE TABLE o_kpi_info(
                           id VARCHAR(32) NOT NULL,
                           name VARCHAR(255),
                           label VARCHAR(255),
                           description VARCHAR(255),
                           unit VARCHAR(255),
                           product_name VARCHAR(255),
                           interval INTEGER,
                           type integer default 0,
                           calc_script TEXT,
                           summary_type VARCHAR(255),
                           value_type VARCHAR(255),
                           org_id VARCHAR(255),
                           PRIMARY KEY (id)
);

COMMENT ON TABLE o_kpi_info IS '指标表';
COMMENT ON COLUMN o_kpi_info.id IS '唯一标识';
COMMENT ON COLUMN o_kpi_info.name IS '英文名';
COMMENT ON COLUMN o_kpi_info.label IS '中文名';
COMMENT ON COLUMN o_kpi_info.description IS '描述';
COMMENT ON COLUMN o_kpi_info.unit IS '单位';
COMMENT ON COLUMN o_kpi_info.product_name IS '计算产品名称';
COMMENT ON COLUMN o_kpi_info.interval IS '分钟粒度 1,5,15,30,60';
COMMENT ON COLUMN o_kpi_info.calc_script IS '计算公式';
COMMENT ON COLUMN o_kpi_info.summary_type IS '统计类型';
COMMENT ON COLUMN o_kpi_info.value_type IS '值类型';
COMMENT ON COLUMN o_kpi_info.org_id IS '组织id';

CREATE TABLE f_kpi_metrics_5m(
                                 id VARCHAR(32) NOT NULL,
                                 device_id VARCHAR(255) NOT NULL,
                                 ts TIMESTAMP NOT NULL,
                                 kpi_id VARCHAR(255) NOT NULL,
                                 value FLOAT4 NOT NULL,
                                 PRIMARY KEY (id,ts)
);

COMMENT ON TABLE f_kpi_metrics_5m IS '5分钟性能表';
COMMENT ON COLUMN f_kpi_metrics_5m.id IS '唯一标识';
COMMENT ON COLUMN f_kpi_metrics_5m.device_id IS '设备id';
COMMENT ON COLUMN f_kpi_metrics_5m.ts IS '时间戳';
COMMENT ON COLUMN f_kpi_metrics_5m.kpi_id IS 'kpi_id';
COMMENT ON COLUMN f_kpi_metrics_5m.value IS '值';

SELECT create_hypertable('f_kpi_metrics_5m','ts');

CREATE MATERIALIZED VIEW f_kpi_metrics_daily
            WITH (timescaledb.continuous)
AS
SELECT
    time_bucket('1 day', ts) as bucket,
    device_id,
    kpi_id,
    avg(value) as avg_value,
    max(value) as max_value,
    sum(value) as sum_value,
    min(value) as min_value

FROM
    f_kpi_metrics_5m
GROUP BY bucket, device_id,kpi_id
WITH NO DATA;


CREATE MATERIALIZED VIEW f_kpi_metrics_monthly
            WITH (timescaledb.continuous)
AS
SELECT
    time_bucket('1 month', ts) as bucket,
    device_id,
    kpi_id,
    avg(value) as avg_value,
    max(value) as max_value,
    sum(value) as sum_value,
    min(value) as min_value

FROM
    f_kpi_metrics_5m
GROUP BY bucket, device_id,kpi_id
WITH NO DATA;


CREATE MATERIALIZED VIEW f_kpi_metrics_yearly
            WITH (timescaledb.continuous)
AS
SELECT
    time_bucket('1 year', ts) as bucket,
    device_id,
    kpi_id,
    avg(value) as avg_value,
    max(value) as max_value,
    sum(value) as sum_value,
    min(value) as min_value

FROM
    f_kpi_metrics_5m
GROUP BY bucket, device_id,kpi_id
WITH NO DATA;




-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

drop view if exists v_device_with_points cascade ;
drop view if exists v_device_tree cascade ;
drop view if exists v_device_identifier_product_json cascade ;
drop view if exists v_device_current_data cascade ;
drop view if exists v_device_point_data cascade ;
drop view if exists v_device_info cascade ;
drop view if exists v_tag_with_product_id cascade ;
drop view if exists v_device_with_tag cascade ;
drop view if exists v_point_with_tag cascade ;
DROP TABLE IF EXISTS o_edge_node cascade;
DROP TABLE IF EXISTS o_user_device cascade;
DROP TABLE IF EXISTS o_device cascade;
DROP TABLE IF EXISTS f_point_data cascade;
DROP TABLE IF EXISTS f_device_data cascade;
DROP TABLE IF EXISTS o_big_screen cascade;

DROP TABLE IF EXISTS o_point_io cascade;
DROP TABLE IF EXISTS o_point_modbus cascade;
DROP TABLE IF EXISTS o_point cascade;
DROP TABLE IF EXISTS o_tag cascade;

DROP TABLE IF EXISTS o_topology cascade;
DROP TABLE IF EXISTS o_group cascade;
DROP TABLE IF EXISTS o_product cascade;
DROP TABLE IF EXISTS o_device_model cascade;
DROP TABLE IF EXISTS o_model_meta cascade;
DROP TABLE IF EXISTS o_alarm_rule cascade;
DROP TABLE IF EXISTS o_device_mirror cascade;
DROP TABLE IF EXISTS o_sim_device cascade;
DROP TABLE IF EXISTS f_sim_device_data cascade;
DROP TABLE IF EXISTS o_project cascade;

DROP TABLE IF EXISTS o_device_group cascade;


DROP TABLE IF EXISTS o_protocol cascade;

DROP TABLE IF EXISTS o_device cascade;

DROP TABLE IF EXISTS o_manage_attribute_meta cascade;

drop table if exists o_topology cascade;

DROP TABLE IF EXISTS o_product cascade;

DROP TABLE IF EXISTS o_kpi_info cascade;
-- +goose StatementEnd

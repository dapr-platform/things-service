dp-cli gen --connstr "postgresql://things:things2024@localhost:5432/thingsdb?sslmode=disable" \
--tables=o_access_protocol,o_device,o_device_mirror,o_point,o_point_io,o_tag,o_sim_device,o_user_device,o_resource,o_alarm_rule,o_product,o_device_model,o_model_meta,o_product_info,o_kpi_info,o_holiday_json \
--model_naming "{{ toUpperCamelCase ( replace . \"o_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"o_\" \"\") }}" \
--module things-service

dp-cli gen --connstr "postgresql://things:things2024@localhost:5432/thingsdb?sslmode=disable" \
--tables=v_device_tree,v_device_current_data,v_product_info,v_device_identifier_product_json,v_tag_with_product_id,v_device_with_tag,v_device_with_tag_filter,v_point_info,v_user_device_info \
--model_naming "{{ toUpperCamelCase ( replace . \"v_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"v_\" \"\") }}" \
--module things-service


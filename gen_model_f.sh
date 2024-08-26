dp-cli gen --connstr "postgresql://things:things2024@localhost:5432/thingsdb?sslmode=disable" \
--tables=f_device_data,f_sim_device_data,f_kpi_metrics_5m \
--model_naming "{{ toUpperCamelCase ( replace . \"f_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"f_\" \"\") }}" \
--module things-service


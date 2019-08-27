#!/usr/bin/env sh

rename() {
    find pkg/client -type d -name "*$1*" | sed "s/\(.*\)$1\(.*\)/mv & \1$2\2/" | sh
    find pkg/client -type f -name "*$1*" | sed "s/\(.*\)$1\(.*\)/mv & \1$2\2/" | sh
}

rename v_sphere vsphere
rename a_w_s aws
rename _d_e_l_e_t_e_ _delete_
rename _g_e_t_ _get_
rename _p_a_t_c_h_ _patch_
rename _p_o_s_t_ _post_
rename _p_u_t_ _put_
rename _g_e_t1_ _get1_

find pkg -type f -exec sed -i '' 's/_a_w_s_/_aws_/g' {} +
find pkg -type f -exec sed -i '' 's/_v_sphere_/_vsphere_/g' {} +
find pkg -type f -exec sed -i '' 's/_d_e_l_e_t_e_/_delete_/g' {} +
find pkg -type f -exec sed -i '' 's/_g_e_t_/_get_/g' {} +
find pkg -type f -exec sed -i '' 's/_p_a_t_c_h_/_patch_/g' {} +
find pkg -type f -exec sed -i '' 's/_p_o_s_t_/_post_/g' {} +
find pkg -type f -exec sed -i '' 's/_p_u_t_/_put_/g' {} +
find pkg -type f -exec sed -i '' 's/ iaa s/ iaas /g' {} +
find pkg -type f -exec sed -i '' 's/ v mware/ vmware/g' {} +
find pkg -type f -exec sed -i '' 's/VMwareCloudAssemblyIaaS/MulticloudIaaS/g' {} +
find pkg -type f -exec sed -i '' 's/ g e t params/ get params/g' {} +
find pkg -type f -exec sed -i '' 's/ g e t 1 / get1 /g' {} +

mv pkg/client/v_mware_cloud_assembly_iaa_s_client.go pkg/client/vra_client.go

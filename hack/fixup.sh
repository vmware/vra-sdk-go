#!/usr/bin/env sh

rename() {
    find pkg/client -type d -name "*$1*" | sed "s/\(.*\)$1\(.*\)/mv & \1$2\2/" | sh
    find pkg/client -type f -name "*$1*" | sed "s/\(.*\)$1\(.*\)/mv & \1$2\2/" | sh
}

rename v_sphere vsphere
rename v_c_f vcf
rename vm_c vmc
rename a_w_s aws
rename _d_e_l_e_t_e_ _delete_
rename _g_e_t_ _get_
rename _p_a_t_c_h_ _patch_
rename _p_o_s_t_ _post_
rename _p_u_t_ _put_
rename _g_e_t1_ _get1_
rename _d_e_l_e_t_e1_ _delete1_
rename _p_a_t_c_h1_ _patch1_
rename _p_o_s_t1_ _post1_
rename _p_u_t1_ _put1_

find pkg -type f -exec sed -i.bak 's/_a_w_s_/_aws_/g' {} +
find pkg -type f -exec sed -i.bak 's/_v_c_f_/_vcf_/g' {} +
find pkg -type f -exec sed -i.bak 's/_v_m_c_/_vmc_/g' {} +
find pkg -type f -exec sed -i.bak 's/_v_sphere_/_vsphere_/g' {} +
find pkg -type f -exec sed -i.bak 's/_d_e_l_e_t_e_/_delete_/g' {} +
find pkg -type f -exec sed -i.bak 's/_g_e_t_/_get_/g' {} +
find pkg -type f -exec sed -i.bak 's/_p_a_t_c_h_/_patch_/g' {} +
find pkg -type f -exec sed -i.bak 's/_p_o_s_t_/_post_/g' {} +
find pkg -type f -exec sed -i.bak 's/_p_u_t_/_put_/g' {} +
find pkg -type f -exec sed -i.bak 's/ iaa s/ iaas /g' {} +
find pkg -type f -exec sed -i.bak 's/ v mware/ vmware/g' {} +
find pkg -type f -exec sed -i.bak 's/VMwareCloudAssemblyIaaS/MulticloudIaaS/g' {} +
find pkg -type f -exec sed -i.bak 's/ g e t params/ get params/g' {} +
find pkg -type f -exec sed -i.bak 's/ g e t 1 / get1 /g' {} +
find pkg -type f -name "*.bak" -exec rm {} +

mv pkg/client/v_mware_cloud_assembly_iaa_s_client.go pkg/client/vra_client.go

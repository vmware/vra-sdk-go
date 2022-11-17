#!/usr/bin/env sh

# Uncomment the other line if backup files are wanted for debugging
BACKUP=""
# BACKUP=".bak"

rename() {
    find pkg/client -type d -name "*$1*" | sed "s/\(.*\)$1\(.*\)/mv & \1$2\2/" | sh
    find pkg/client -type f -name "*$1*" | sed "s/\(.*\)$1\(.*\)/mv & \1$2\2/" | sh
}

rename v_sphere vsphere
rename v_center vcenter
rename v_c_f vcf
rename vm_c vmc
rename a_w_s aws
rename p_k_s pks
rename _d_e_l_e_t_e _delete
rename _g_e_t _get
rename _p_a_t_c_h _patch
rename _p_o_s_t _post
rename _p_u_t _put
rename _g_e_t1 _get1
rename _d_e_l_e_t_e1 _delete1
rename _p_a_t_c_h1 _patch1
rename _p_o_s_t1 _post1
rename _p_u_t1 _put1

find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/_a_w_s_/_aws_/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/_v_c_f_/_vcf_/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/v_c_f/vcf/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/_v_m_c_/_vmc_/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/_v_sphere_/_vsphere_/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/v_sphere_/vsphere_/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/v_center_/vcenter_/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/p_k_s_/pks_/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/_d_e_l_e_t_e/_delete/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/_g_e_t/_get/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/_p_a_t_c_h/_patch/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/_p_o_s_t/_post/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/_p_u_t/_put/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/ iaa s/ iaas /g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/ v mware/ vmware/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/VMwareCloudAssemblyBlueprintAPI/API/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/ g e t params/ get params/g' {} +
find pkg -type f -name '*.go' -exec sed -i "$BACKUP" 's/ g e t 1 / get1 /g' {} +

mv pkg/client/v_mware_cloud_assembly_blueprint_api_client.go pkg/client/vra_client.go
sed -i "$BACKUP" 's/^\(.*DefaultHost string = \).*/\1"api.mgmt.cloud.vmware.com"/' ./pkg/client/vra_client.go

# cleanup if needed
find pkg -type f -name "*.bak" -exec rm {} +

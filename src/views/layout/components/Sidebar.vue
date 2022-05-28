<template>
  <div class="sidebar-container">
    <el-menu
      background-color="#001529"
      text-color="#fff"
      active-text-color="#409eff"
      :collapse="isCollapse"
      :collapse-transition="false"
      unique-opened
      router
      :default-active="$store.state.user.activePath"
    >
      <!-- S 一级菜单 -->
      <template v-for="item in menuList" :key="item.path">
        <el-sub-menu :index="item.path + ''" v-if="item.children && item.children.length">
          <template #title>
            <span>{{ item.permissionName }}</span>
          </template>

          <!-- S 二级菜单 -->
          <template v-for="itemChild in item.children" :key="itemChild.path">
            <el-sub-menu
              :index="itemChild.path + 'tool'"
              v-if="itemChild.children && itemChild.children.length"
            >
              <template #title>
                <span>{{ itemChild.permissionName }}</span>
              </template>

              <!-- S 三级菜单-->
              <template v-for="itemChild_child in itemChild.children" :key="itemChild_child.path">
                <el-sub-menu
                  :index="itemChild_child.path + ''"
                  v-if="itemChild_child.children && itemChild_child.children.length"
                >
                  <template #title>
                    <span>{{ itemChild_child.permissionName }}</span>
                  </template>

                  <!-- S 四级菜单-->
                  <template
                    v-for="itemChild_child_th in itemChild_child.children"
                    :key="itemChild_child_th.path"
                  >
                    <el-menu-item :index="itemChild_child_th.path">
                      <a
                        v-if="!itemChild_child_th.onSite"
                        :href="itemChild_child_th.path"
                        target="_blank"
                      >
                        {{ itemChild_child_th.permissionName }}
                      </a>
                      <template #title>
                        <span>
                          {{ itemChild_child_th.permissionName }}
                        </span>
                      </template>
                    </el-menu-item>
                  </template>
                  <!-- E 四级菜单-->
                </el-sub-menu>
                <el-menu-item v-else :index="itemChild_child.path" :key="itemChild_child.path">
                  <a v-if="!itemChild_child.onSite" :href="itemChild_child.path" target="_blank">
                    {{ itemChild_child.permissionName }}
                  </a>
                  <template v-if="itemChild_child.onSite" #title>
                    <span>{{ itemChild_child.permissionName }}</span>
                  </template>
                </el-menu-item>
              </template>
              <!-- E 三级菜单-->
            </el-sub-menu>
            <el-menu-item
              v-else
              :index="!itemChild.onSite ? '' : itemChild.path"
              :key="itemChild.path"
            >
              <a v-if="!itemChild.onSite" :href="itemChild.path" target="_blank">
                {{ itemChild.permissionName }}
              </a>
              <template v-if="itemChild.onSite" #title>
                <span>
                  {{ itemChild.permissionName }}
                </span>
              </template>
            </el-menu-item>
          </template>
          <!-- E 二级菜单 -->
        </el-sub-menu>
        <el-menu-item v-else :index="item.path" :key="item.path">
          <template v-if="!item.onSite">
            <a :href="item.path" target="_blank">
              {{ item.permissionName }}
            </a>
          </template>
          <template v-if="item.onSite" #title>
            <span>{{ item.permissionName }}</span>
          </template>
        </el-menu-item>
      </template>
      <!-- E 一级菜单 -->
    </el-menu>
  </div>
</template>

<script>
export default {
  name: "CommonSidebar",
  props: {},
  data() {
    return {
      menuList: [],
      // 是否显示悬浮菜单
      isCollapse: false,
      menuList2: [
        {
          "permissionId": 105,
          "enableFlag": true,
          "permissionName": "教学中心",
          "onSite": true,
          "path": "customer",
          "permissionTypeId": 1,
          "orderNum": 0,
          "children": [
            {
              "permissionId": 106,
              "enableFlag": true,
              "parentId": 105,
              "permissionName": "课程信息",
              "onSite": true,
              "path": "/crm/customer",
              "permissionTypeId": 2,
              "orderNum": 0
            }
          ]
        }
      ]
    };
  },
  computed: {},
  watch: {},
  created() {
    //获取菜单列表
    this.getMenuList();
  },
  methods: {
    // 获取菜单列表
    getMenuList() {
      this.menuList = this.menuList2;
    }
  }
};
</script>
<style lang="less" scoped>
.sidebar-container {
  height: 100%;
  background-color: #001529;
  padding-top: 60px;
}
</style>

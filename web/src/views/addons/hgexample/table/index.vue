<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="表格例子">
        这里提供了一些常用的普通表格组件的用法和表单组件的例子，你可能会需要
      </n-card>
    </div>
    <n-card :bordered="false" class="proCard">
      <BasicForm
        @register="register"
        @submit="reloadTable"
        @reset="reloadTable"
        @keyup.enter="reloadTable"
        ref="searchFormRef"
      >
        <template #statusSlot="{ model, field }">
          <n-input v-model:value="model[field]" />
        </template>
      </BasicForm>

      <BasicTable
        :openChecked="true"
        :columns="columns"
        :request="loadDataTable"
        :row-key="(row) => row.id"
        ref="actionRef"
        :actionColumn="actionColumn"
        :checked-row-keys="checkedIds"
        @update:checked-row-keys="onCheckedRow"
        :scroll-x="scrollX"
        :resizeHeightOffset="-10000"
        size="small"
        @update:sorter="handleUpdateSorter"
      >
        <template #tableTitle>
          <n-button type="primary" @click="addTable" class="min-left-space">
            <template #icon>
              <n-icon>
                <PlusOutlined />
              </n-icon>
            </template>
            添加
          </n-button>
          <n-button
            type="error"
            @click="handleBatchDelete"
            :disabled="batchDeleteDisabled"
            class="min-left-space"
          >
            <template #icon>
              <n-icon>
                <DeleteOutlined />
              </n-icon>
            </template>
            批量删除
          </n-button>
          <n-button type="primary" @click="handleExport" class="min-left-space">
            <template #icon>
              <n-icon>
                <ExportOutlined />
              </n-icon>
            </template>
            导出
          </n-button>
        </template>
      </BasicTable>
    </n-card>
    <Edit
      @reload-table="reloadTable"
      @update-show-modal="updateShowModal"
      :showModal="showModal"
      :formParams="formParams"
    />
  </div>
</template>

<script lang="ts" setup>
  import { computed, h, onMounted, reactive, ref } from 'vue';
  import { useDialog, useMessage } from 'naive-ui';
  import { BasicTable, TableAction } from '@/components/Table';
  import { BasicForm, useForm } from '@/components/Form/index';
  import { useSorter } from '@/hooks/common';
  import { Delete, List, Status, Export } from '@/api/addons/hgexample/table';
  import { State, columns, schemas, newState, loadOptions } from './model';
  import { DeleteOutlined, PlusOutlined, ExportOutlined } from '@vicons/antd';
  import { useRouter } from 'vue-router';
  import { adaTableScrollX } from '@/utils/hotgo';
  import Edit from './edit.vue';
  import { useDictStore } from '@/store/modules/dict';

  const dict = useDictStore();
  const router = useRouter();
  const dialog = useDialog();
  const message = useMessage();
  const searchFormRef = ref<any>();
  const batchDeleteDisabled = ref(true);
  const checkedIds = ref([]);
  const showModal = ref(false);
  const formParams = ref<State>();
  const actionRef = ref();
  const { updateSorter: handleUpdateSorter, sortStatesRef: sortStatesRef } = useSorter(reloadTable);

  const actionColumn = reactive({
    width: 300,
    title: '操作',
    key: 'action',
    fixed: 'right',
    render(record) {
      return h(TableAction as any, {
        style: 'button',
        actions: [
          {
            label: '编辑',
            onClick: handleEdit.bind(null, record),
          },
          {
            label: '禁用',
            onClick: handleStatus.bind(null, record, 2),
            ifShow: () => {
              return record.status === 1;
            },
          },
          {
            label: '启用',
            onClick: handleStatus.bind(null, record, 1),
            ifShow: () => {
              return record.status === 2;
            },
          },
          {
            label: '删除',
            onClick: handleDelete.bind(null, record),
          },
        ],
        dropDownActions: [
          {
            label: '查看详情',
            key: 'view',
          },
          {
            label: '更多按钮1',
            key: 'test1',
            // 根据业务控制是否显示: 非enable状态的不显示启用按钮
            ifShow: () => {
              return true;
            },
          },
          {
            label: '更多按钮2',
            key: 'test2',
            ifShow: () => {
              return true;
            },
          },
        ],
        select: (key) => {
          if (key === 'view') {
            return handleView(record);
          }
          message.info(`您点击了，${key} 按钮`);
        },
      });
    },
  });

  const scrollX = computed(() => {
    return adaTableScrollX(columns, actionColumn.width);
  });

  const [register, {}] = useForm({
    gridProps: { cols: '1 s:1 m:2 l:3 xl:4 2xl:4' },
    labelWidth: 80,
    schemas,
  });

  const loadDataTable = async (res) => {
    return await List({
      ...searchFormRef.value?.formModel,
      ...{ sorters: sortStatesRef.value },
      ...res,
    });
  };

  function addTable() {
    showModal.value = true;
    formParams.value = newState(null);
  }

  function updateShowModal(value) {
    showModal.value = value;
  }

  function onCheckedRow(rowKeys) {
    batchDeleteDisabled.value = rowKeys.length <= 0;
    checkedIds.value = rowKeys;
  }

  function reloadTable() {
    actionRef.value.reload();
  }

  function handleView(record: Recordable) {
    router.push({ name: 'table_view', params: { id: record.id } });
  }

  function handleEdit(record: Recordable) {
    showModal.value = true;
    formParams.value = newState(record as State);
  }

  function handleDelete(record: Recordable) {
    dialog.warning({
      title: '警告',
      content: '你确定要删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete(record).then((_res) => {
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  function handleBatchDelete() {
    dialog.warning({
      title: '警告',
      content: '你确定要批量删除？',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Delete({ id: checkedIds.value }).then((_res) => {
          batchDeleteDisabled.value = true;
          checkedIds.value = [];
          message.success('删除成功');
          reloadTable();
        });
      },
    });
  }

  function handleExport() {
    message.loading('正在导出列表...', { duration: 1200 });
    Export(searchFormRef.value?.formModel);
  }

  function handleStatus(record: Recordable, status: number) {
    Status({ id: record.id, status: status }).then((_res) => {
      message.success('设为' + dict.getLabel('sys_normal_disable', status) + '成功');
      setTimeout(() => {
        reloadTable();
      });
    });
  }

  onMounted(() => {
    loadOptions();
  });
</script>

<style lang="less" scoped></style>

<script setup>
import { ref } from 'vue'
import Shares from '@/utils/server/shares'

const tableData = ref([])
const activeName = ref(1)
const loading = ref(false)
const tableRef = ref()
const tabs = ref([
    {
        label: '',
        value: 1
    },
    {
        label: '近5日',
        value: 5
    },
    {
        label: '近10日',
        value: 10
    },
    {
        label: '近20日',
        value: 20
    },
    {
        label: '近100日',
        value: 100
    }
])
const emit = defineEmits(['ok'])
const init = async (call) => {
    loading.value = true
    try {
        const res = await Shares.GetHyZyb({ tag: 6, days: activeName.value })
        loading.value = false
        if (res.list) {
            tableData.value = res.list
            tabs.value[0].label = res.day0
            call(res.list[0]?.code || '')
            tableRef.value.setCurrentRow(res.list[0])
        }
    } catch (error) {
        loading.value = false
    }
}

const ok = (row) => emit('ok', row.code)

defineExpose({
    init,
})
</script>
<template>
    <el-tabs v-model="activeName" @tab-change="init">
        <el-tab-pane v-for="i in tabs" :key="i.value" :label="i.label" :name="i.value" />
    </el-tabs>
    <el-table ref="tableRef" highlight-current-row v-loading="loading" :data="tableData" style="width: 100%" height="100%"
        @row-click="ok">
        <el-table-column prop="name" label="股票名称">
            <template #default="{ row }">
                <span class="fontSize" style="background-color: rgb(235, 90, 102);color:#fff;">{{ row.ext }}</span> {{
                    row.name }}
            </template>
        </el-table-column>
        <el-table-column prop="sPrice" label="平均换手率">
            <template #default="{ row }">
                <span :style="`color:${row.img}`">{{ row.sPrice }}</span>
            </template>
        </el-table-column>
        <el-table-column prop="percent" label="涨跌幅">
            <template #default="{ row }">
                <span :style="`color:${row.color}`">{{ row.percent.toFixed(2) }}%</span>
            </template>
        </el-table-column>
    </el-table>
</template>

<style lang="scss">
.fontSize {
    font-size: 12px;
    padding: 1px;
}

.el-tabs__item:hover,
.el-tabs__item.is-active {
    color: #aa55ff;
}
.el-tabs__active-bar {
    background-color: #aa55ff;
}
.el-table__body tr.current-row>td.el-table__cell {
    background-color: #aa55ff44;
}
</style>
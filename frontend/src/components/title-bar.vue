<script setup lang="ts">
import { isDark, toggleDark } from 'vue-dark-switch'
import logo from '~/assets/imgs/logo.png'
import { NRadio } from 'naive-ui'
import emitter from '~/tools/emitter'
import {
	MENU_OPEN_DIR,
	MENU_OPEN_URL,
	MENU_OPEN_FILE,
	MENU_EXPORT,
} from '~/constants/evt-name'
import setting from '~/components/setting.vue'

const keys = useMagicKeys()

const menu = [
	{
		label: '文件(F)',
		key: '文件(F)',
		submenu: [
			{
				label: '打开文件...',
				key: '打开文件...',
				accelerator: 'Ctrl+O',
				onClick: () => {
					emitter.emit(MENU_OPEN_FILE)
				},
			},
			{
				label: '打开文件夹...',
				key: '打开文件夹...',
				accelerator: '',
				onClick: () => {
					emitter.emit(MENU_OPEN_DIR)
				},
			},
			{
				label: '打开网络地址...',
				key: '打开网络地址...',
				accelerator: '',
				onClick: () => {
					emitter.emit(MENU_OPEN_URL)
				},
			},
			{
				label: '导出',
				key: '导出',
				accelerator: '',
				onClick: () => {
					emitter.emit(MENU_EXPORT)
				},
			},
			{
				label: '设置',
				key: '设置',
				accelerator: 'Ctrl+.',
				onClick: () => {
					state.isSettingVisible = true
				},
			},
			{ type: 'divider' },
			{
				label: '退出',
				key: '退出',
				accelerator: '',
				onClick: () => {
					runtime.Quit()
				},
			},
		],
	},
	{
		label: '查看(V)',
		key: '查看(V)',
		submenu: [
			{
				label: '刷新',
				key: '刷新',
				accelerator: 'F5',
				onClick: () => {
					location.reload()
				},
			},
			// { type: 'divider' },
			// {
			// 	label: '实际大小',
			// 	key: '实际大小',
			// 	accelerator: 'Ctrl+0',
			// 	onClick: () => {},
			// },
			// {
			// 	label: '放大',
			// 	key: '放大',
			// 	accelerator: 'Ctrl+=',
			// 	onClick: () => {},
			// },
			// {
			// 	label: '缩小',
			// 	key: '缩小',
			// 	accelerator: 'Ctrl+Minus',
			// 	acceleratorAlias: 'Ctrl+-',
			// 	onClick: () => {},
			// },
			// { type: 'divider' },
			// {
			// 	label: '开发者工具',
			// 	key: '开发者工具',
			// 	accelerator: 'F12',
			// 	onClick: () => {},
			// },
		],
	},
	{
		label: '主题(T)',
		key: '主题(T)',
		submenu: [
			{
				label: '浅色主题',
				key: '浅色主题',
				isChecked: computed(() => !isDark.value),
				onClick: () => {
					isDark.value = false
				},
			},
			{
				label: '深色主题',
				key: '深色主题',
				isChecked: computed(() => isDark.value),
				onClick: () => {
					isDark.value = true
				},
			},
		],
	},
	{
		label: '帮助(H)',
		key: '帮助(H)',
		submenu: [
			{
				label: '官网',
				key: '官网',
				onClick: () => {
					runtime.BrowserOpenURL('https://dzishop.sumoli.com')
				},
			},
			{
				label: 'Github',
				key: 'Github',
				onClick: () => {
					runtime.BrowserOpenURL('https://github.com/WT-SML/dzishop-go')
				},
			},
			{
				label: '更新日志',
				key: '更新日志',
				onClick: () => {},
			},
			{ type: 'divider' },
			{
				label: '检查更新...',
				key: '检查更新...',
				onClick: () => {},
			},
			{
				label: '关于',
				key: '关于',
				onClick: () => {},
			},
		],
	},
]

// 处理下拉菜单
const handleMenuSelect = (key, option) => {
	option?.onClick()
}
// 下拉菜单渲染
const renderLabel = (option) => {
	return h(
		'div',
		{
			className: 'flex justify-between',
		},
		[
			option.isChecked !== undefined
				? h(
						NRadio,
						{
							checked: option.isChecked,
						},
						option.label,
					)
				: h('span', option.label),
			option.accelerator
				? h(
						'span',
						{ className: 'ml-10' },
						option.acceleratorAlias || option.accelerator,
					)
				: null,
		],
	)
}
// @ts-ignore
const runtime = window.runtime
// 监听快捷键
const listenShortcut = () => {
	const menuItems: any = menu.map((item) => item.submenu).flat()
	for (const v of menuItems) {
		if (!v.accelerator) continue
		whenever(keys[v.accelerator], () => {
			v.onClick()
		})
	}
}

const state = reactive({
	isSettingVisible: false,
})

onMounted(() => {
	listenShortcut()
})
</script>

<template>
	<div
		:class="`h-35px bg-[#${isDark ? '21252B' : 'F8F8F8'}] flex justify-between select-none titlebar`"
		style="--wails-draggable: drag"
		@dblclick="runtime.WindowToggleMaximise"
	>
		<div class="h-full flex items-center pl-3">
			<img :src="logo" class="mb-2px mr-5 w-18px" />
			<div>
				<n-dropdown
					v-for="(item, i) in menu"
					:key="item.key"
					:options="item.submenu"
					trigger="click"
					placement="bottom-start"
					:animated="false"
					:render-label="renderLabel"
					@select="handleMenuSelect"
				>
					<div
						:class="`menu-item hover:${isDark ? 'bg-[#34383D]' : 'bg-[#E4E4E4]'}`"
					>
						{{ item.label }}
					</div>
				</n-dropdown>
			</div>
		</div>
		<div :class="`text-14px ${isDark ? 'text-[#C5C5C5]' : 'text-[#333639]'}`">
			<div
				:class="`titlebar-button hover:${isDark ? 'bg-[#40444A]' : 'bg-[#D7D7D7]'}`"
				@click="runtime.WindowMinimise()"
			>
				<div :class="`i-mdi:window-minimize `"></div>
			</div>
			<div
				:class="`titlebar-button hover:${isDark ? 'bg-[#40444A]' : 'bg-[#D7D7D7]'}`"
				@click="runtime.WindowToggleMaximise()"
			>
				<div :class="`i-mdi:window-maximize `"></div>
			</div>
			<div
				class="titlebar-button hover:bg-[#E81123] hover:text-[#fff]"
				@click="runtime.Quit()"
			>
				<div :class="`i-mdi:window-close`"></div>
			</div>
		</div>
	</div>
	<setting
		:isVisible="state.isSettingVisible"
		@close="state.isSettingVisible = false"
	/>
</template>

<style lang="scss" scoepd>
.titlebar {
	.menu-item {
		margin-right: 5px;
		display: inline-block;
		padding: 3px 6px;
		border-radius: 4px;
		cursor: pointer;
	}
	.titlebar-button {
		display: inline-flex;
		justify-content: center;
		align-items: center;
		width: 35px;
		height: 35px;
		cursor: pointer;
	}
}
</style>

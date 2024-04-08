<script setup lang="ts">
// @ts-nocheck
import { isDark, toggleDark } from 'vue-dark-switch'
import logo from '~/assets/imgs/logo.png'
import { appWindow } from '@tauri-apps/api/window'
import { NRadio } from 'naive-ui'

const menu = [
	{
		label: '文件(F)',
		key: '文件(F)',
		submenu: [
			{
				label: '打开文件...',
				key: '打开文件...',
				accelerator: 'Ctrl+O',
				onClick: () => {},
			},
			{
				label: '打开文件夹...',
				key: '打开文件夹...',
				accelerator: '',
				onClick: () => {},
			},
			{
				label: '打开网络地址...',
				key: '打开网络地址...',
				accelerator: '',
				onClick: () => {},
			},
			{
				label: '导出',
				key: '导出',
				accelerator: '',
				onClick: () => {},
			},
			{
				label: '设置',
				key: '设置',
				accelerator: 'Ctrl+.',
				onClick: () => {},
			},
			{ type: 'divider' },
			{
				label: '退出',
				key: '退出',
				accelerator: '',
				onClick: () => {
					appWindow.close()
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
				onClick: () => {},
			},
			{
				label: '强制刷新',
				key: '强制刷新',
				accelerator: 'Ctrl+F5',
				onClick: () => {},
			},

			{ type: 'divider' },
			{
				label: '切换全屏',
				key: '切换全屏',
				accelerator: 'F11',
				onClick: () => {},
			},
			{ type: 'divider' },
			{
				label: '实际大小',
				key: '实际大小',
				accelerator: 'Ctrl+0',
				role: 'resetzoom',
				onClick: () => {},
			},
			{
				label: '放大',
				key: '放大',
				accelerator: 'Ctrl+=',
				role: 'zoomin',
				onClick: () => {},
			},
			{
				label: '缩小',
				key: '缩小',
				accelerator: 'Ctrl+-',
				role: 'zoomout',
				onClick: () => {},
			},

			{ type: 'divider' },
			{
				label: '开发者工具',
				key: '开发者工具',
				role: 'toggledevtools',
				accelerator: 'F12',
				onClick: () => {},
			},
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
				onClick: () => {},
			},
			{
				label: 'Github',
				key: 'Github',
				onClick: () => {},
			},
			{
				label: '更新日志',
				key: '更新日志',
				onClick: () => {},
			},
			{
				label: '隐私条款',
				key: '隐私条款',
				onClick: () => {},
			},
			{ type: 'divider' },
			{
				label: '检查更新...',
				key: '检查更新...',
				onClick: () => {},
			},
			{
				label: '查看许可证...',
				key: '查看许可证...',
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
				? h('span', { className: 'ml-10' }, option.accelerator)
				: null,
		],
	)
}

const runtime = window.runtime
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
		<div :class="`${isDark ? 'text-[#C5C5C5]' : 'text-[#333639]'}`">
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
</template>

<style lang="scss" scoepd>
.titlebar {
	.menu-item {
		margin-right: 5px;
		display: inline-block;
		padding: 3px 6px;
		border-radius: 4px;
	}
	.titlebar-button {
		display: inline-flex;
		justify-content: center;
		align-items: center;
		width: 35px;
		height: 35px;
	}
}
</style>

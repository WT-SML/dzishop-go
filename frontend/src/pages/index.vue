<script setup lang="ts">
import { isDark } from 'vue-dark-switch'
import _ from 'lodash'
import { SUPPORT_FILE_TYPES } from '@/constants'
import { sortByKey, p, isChineseCharacter, isNumber } from '@/tools'
import { useMagicKeys } from '@vueuse/core'
import {
	WailsGetDirectoryItem,
	WailsSelectFolder,
} from '../../wailsjs/go/main/App'
import emitter from '~/tools/emitter'
import {
	MENU_OPEN_FILE,
	MENU_OPEN_DIR,
	MENU_OPEN_URL,
} from '~/constants/evt-name'
import { pinyin } from 'pinyin-pro'

// 打开文件
emitter.on(MENU_OPEN_FILE, () => {
	// for (const filePath of filePaths) {
	// 	if (state.tabs.map((item) => item.key).includes(filePath)) {
	// 		state.activeTabKey = filePath
	// 		syncActiveTabToRsourceManager()
	// 		return
	// 	}
	// 	state.tabs.push({
	// 		key: filePath,
	// 		label: path.basename(filePath),
	// 	})
	// 	state.activeTabKey = filePath
	// 	syncActiveTabToRsourceManager()
	// }
})
// 打开文件夹
emitter.on(MENU_OPEN_DIR, () => {
	openFolder()
})
// 打开网络地址
emitter.on(MENU_OPEN_URL, () => {
	state.websiteAddressModal.websiteAddress = null
	state.websiteAddressModal.isModalShow = true
})

// 拖放的Dom引用
const tabsRef = ref(null)

// 组合键
const keys = useMagicKeys()
const CtrlW = keys['Ctrl+W']
watch(CtrlW, (v) => {
	if (v) {
		if (state.activeTabKey) {
			closeTab(state.tabs.map((item) => item.key).indexOf(state.activeTabKey))
		}
	}
})

// 菜单
const lightThemeColor = '#f8f8f8'
const darkThemeColor = '#282c34'
const frameSizeSettingRef = ref(null)
// 资源管理器面板宽度
const compotedResourceManagerPanelWidth = computed(() => {
	return `${state.resourceManagerPanelWidth}px`
})
// 全局鼠标的位置
const { x: mousePositionX, y: mousePositionY } = useMouse()
const { pressed: isFrameSizeSettingRefMousePressed } = useMousePressed({
	target: frameSizeSettingRef,
})
let mousePositionXWatchStop = null
watch(isFrameSizeSettingRefMousePressed, (newVal) => {
	if (newVal) {
		let oldWidth = state.resourceManagerPanelWidth
		let oldMousePositionX = mousePositionX.value
		mousePositionXWatchStop = watch(mousePositionX, (newMousePositionX) => {
			let newWidth = 0
			newWidth = oldWidth + (newMousePositionX - oldMousePositionX)
			if (newWidth < 150) {
				newWidth = 150
			} else if (newWidth > 500) {
				newWidth = 500
			}
			state.resourceManagerPanelWidth = newWidth
		})
	} else {
		mousePositionXWatchStop && mousePositionXWatchStop()
	}
})
// 状态
const defaultState = {
	resourceManagerPanelWidth: 250, // 资源管理器的面板宽度
	folderList: [], // 文件夹列表
	defaultExpandedKeys: [], // 默认展开项
	selectedKeys: [], // 选中节点
	tabs: [], //标签页列表
	activeTabKey: null, //激活的tab的key
	realTimeCoord: null, // 实际图像坐标
	// 网络地址模态框
	websiteAddressModal: {
		isModalShow: false,
		websiteAddress: null,
	},
}
let initState
if (localStorage.getItem('state')) {
	initState = JSON.parse(localStorage.getItem('state'))
} else {
	initState = defaultState
}

const state = reactive(initState)
// 监听需要缓存的项
watch(
	() => state,
	_.debounce(function (newVal) {
		localStorage.setItem('state', JSON.stringify(state))
	}, 500),
	{
		deep: true,
	},
)
// 提交网络地址
const websiteAddressSubmit = () => {
	if (!state.websiteAddressModal.websiteAddress) {
		return
	}
	const websiteAddress = state.websiteAddressModal.websiteAddress
	if (state.tabs.map((item) => item.key).includes(websiteAddress)) {
		state.activeTabKey = websiteAddress
		syncActiveTabToRsourceManager()
		state.websiteAddressModal.isModalShow = false
		return
	}
	const label = websiteAddress.split('/').at(-1).slice(0, -4)
	state.tabs.push({
		key: websiteAddress,
		label,
		isWebsiteAddress: true,
	})
	state.activeTabKey = websiteAddress
	syncActiveTabToRsourceManager()
	state.websiteAddressModal.isModalShow = false
}
// 点击打开文件夹按钮
const openFolder = async () => {
	const folderPath = await WailsSelectFolder()
	if (!folderPath) return
	const ret = await getDirectoryItem(folderPath)
	state.folderList[0] = ret
	state.defaultExpandedKeys = [ret.key]
}
// 获取一个目录下的项目
const getDirectoryItem = async (dirPath) => {
	let entries = p(await WailsGetDirectoryItem(dirPath))
	for (const v of entries) {
		const firstChar = v.label[0].toLocaleLowerCase()
		v.sortField = isChineseCharacter(firstChar)
			? pinyin(firstChar)[0]
			: firstChar
		if (!v.isLeaf) {
			v.children = null
		}
	}
	// 排序
	// 文件夹
	const dirs = entries.filter((item) => !item.isLeaf)
	// 文件夹-数字
	const dirsNumber = dirs.filter((item) => isNumber(item.label[0]))
	sortByKey(dirsNumber, 'sortField')
	// 文件夹-汉字
	const dirsChinese = dirs.filter((item) => isChineseCharacter(item.label[0]))
	sortByKey(dirsChinese, 'sortField')
	// 文件夹-非数字汉字
	const dirsNotChineseOrNumber = dirs.filter(
		(item) => !isChineseCharacter(item.label[0]) && !isNumber(item.label[0]),
	)
	sortByKey(dirsNotChineseOrNumber, 'sortField')
	// 文件
	const files = entries.filter((item) => item.isLeaf)
	// 文件-数字
	const filesNumber = files.filter((item) => isNumber(item.label[0]))
	sortByKey(filesNumber, 'sortField')
	// 文件-汉字
	const filesChinese = files.filter((item) => isChineseCharacter(item.label[0]))
	sortByKey(filesChinese, 'sortField')
	// 文件夹-非数字汉字
	const filesNotChineseOrNumber = files.filter(
		(item) => !isChineseCharacter(item.label[0]) && !isNumber(item.label[0]),
	)
	sortByKey(filesNotChineseOrNumber, 'sortField')
	entries = [
		// 文件夹
		...[
			// 文件夹-数字
			...dirsNumber,
			// 文件夹-汉字
			...dirsChinese,
			// 文件夹-非数字汉字
			...dirsNotChineseOrNumber,
		],
		...[
			// 文件-数字
			...filesNumber,
			// 文件-汉字
			...filesChinese,
			// 文件-非汉字
			...filesNotChineseOrNumber,
		],
	]
	const label = dirPath.split('\\').at(-1)
	const ret = {
		label,
		key: dirPath,
		children: entries,
		isLeaf: false,
	}
	return ret
}
// label渲染
const renderLabel = ({ option }) => {
	return h(
		'div',
		{
			title: option.key,
			style: {
				overflow: 'hidden',
				'white-space': 'nowrap',
				'text-overflow': 'ellipsis',
			},
		},
		`${option.label}`,
	)
}
// 前置图标渲染
const renderPrefix = ({ option }) => {
	return !option.isLeaf
		? null
		: h('div', {
				className: 'i-ic:round-insert-drive-file ml--11px',
			})
}
// 处理节点加载
const handleNodeLoad = (node) => {
	return new Promise<void>((resolve) => {
		setTimeout(async () => {
			node.children = (await getDirectoryItem(node.key)).children
			resolve()
		}, 0)
	})
}
// 处理节点选中项发生变化
const handleSelectedKeysUpdate = (keys, option, meta) => {
	const target = option[0]
	const key = target.key
	if (!target.isLeaf) {
		return
	}
	if (state.tabs.map((item) => item.key).includes(key)) {
		state.activeTabKey = key
		return
	}
	state.tabs.push({
		key,
		label: target.label,
	})
	state.activeTabKey = key
}
// 关闭tab
const closeTab = (i) => {
	let activeIndex = state.tabs
		.map((item) => item.key)
		.indexOf(state.activeTabKey)
	if (i === activeIndex) {
		if (i === state.tabs.length - 1) {
			activeIndex -= 1
		}
	} else if (i < activeIndex) {
		activeIndex -= 1
	}
	state.tabs.splice(i, 1)
	state.activeTabKey = activeIndex < 0 ? null : state.tabs[activeIndex].key
	syncActiveTabToRsourceManager()
}
// 处理标签页切换
const handleTabChange = (i) => {
	const item = state.tabs[i]
	state.activeTabKey = item.key
	syncActiveTabToRsourceManager()
}
// 激活选项同步到左侧
const syncActiveTabToRsourceManager = () => {
	state.selectedKeys = state.activeTabKey === null ? [] : [state.activeTabKey]
}
// 挂载
onMounted(async () => {
	// 阻止浏览器默认的拖放行为
	tabsRef.value.addEventListener('dragover', (e) => {
		e.preventDefault()
	})
	// 处理文件拖放事件
	tabsRef.value.addEventListener('drop', (e) => {
		e.preventDefault()
		const files = e.dataTransfer.files
		// 处理拖放的文件
		// for (const file of files) {
		// 	if (file.path) {
		// 		if (state.tabs.map((item) => item.key).includes(file.path)) {
		// 			state.activeTabKey = file.path
		// 			syncActiveTabToRsourceManager()
		// 			return
		// 		}
		// 		state.tabs.push({
		// 			key: file.path,
		// 			label: path.basename(file.path),
		// 		})
		// 		state.activeTabKey = file.path
		// 		syncActiveTabToRsourceManager()
		// 	}
		// }
	})
})
</script>
<template>
	<div
		class="main w-[100vw] flex flex-col overflow-hidden"
		:style="{
			userSelect: isFrameSizeSettingRefMousePressed ? 'none' : 'unset',
			cursor: isFrameSizeSettingRefMousePressed ? 'e-resize' : '',
		}"
	>
		<!-- 主要部分 -->
		<div class="flex flex-grow overflow-hidden">
			<!-- 资源管理器 -->
			<div
				:class="`resource-manager h-full max-w-[600px] flex-shrink-0 bg-[#${
					isDark ? '21252B' : 'F8F8F8'
				}] relative flex flex-col select-none ${
					isDark ? '' : 'border-t border-[#dfdfdf]'
				}`"
				:style="{ width: compotedResourceManagerPanelWidth }"
			>
				<div
					:class="`font-bold h-30px flex items-center flex-shrink-0 px-20px ${
						isDark ? '' : 'border-b border-[#dfdfdf]'
					}`"
				>
					资源管理器
				</div>
				<div v-if="!state.folderList.length" class="px-20px">
					<div class="mb-3">尚未打开文件夹。</div>
					<n-button type="primary" block @click="openFolder()">
						打开文件夹
					</n-button>
				</div>
				<div v-else class="flex-grow overflow-x-hidden overflow-y-auto pb-40px">
					<n-tree
						v-model:default-expanded-keys="state.defaultExpandedKeys"
						v-model:selected-keys="state.selectedKeys"
						:on-load="handleNodeLoad"
						expand-on-click
						block-line
						block-node
						:render-prefix="renderPrefix"
						:render-label="renderLabel"
						:data="state.folderList"
						:animated="false"
						:cancelable="false"
						@update:selected-keys="handleSelectedKeysUpdate"
					/>
				</div>
				<!-- 窗口大小调整器 -->
				<div
					ref="frameSizeSettingRef"
					class="frame-size-setting absolute right-[-3px] top-0 z-1 h-full w-4px select-none hover:bg-[#005FB8]"
					:style="{
						backgroundColor: isFrameSizeSettingRefMousePressed ? '#005FB8' : '',
					}"
				></div>
			</div>
			<!-- tabs -->
			<div
				ref="tabsRef"
				:class="`tabs flex-grow h-full select-none overflow-hidden ${
					isDark ? '' : 'border-t border-l border-[#dfdfdf]'
				}`"
				:style="{ backgroundColor: isDark ? '#23272E' : '#fff' }"
			>
				<div
					v-if="!state.tabs.length"
					class="h-full w-full flex flex-col items-center justify-center"
				>
					<div class="mb-3">一款优秀的深度缩放图像查看器</div>
					<div>支持文件格式：{{ SUPPORT_FILE_TYPES.join(' | ') }}</div>
				</div>
				<div v-else class="h-full w-full flex flex-col">
					<div
						:class="`tabs-bar flex-shrink-0 flex items-center h-30px overflow-hidden bg-[#${
							isDark ? '21252B' : 'F8F8F8'
						}] ${isDark ? '' : 'border-b border-[#dfdfdf]'}`"
					>
						<div
							v-for="(item, i) in state.tabs"
							:key="item.key"
							:class="[
								`h-full flex items-center justify-between pl-12px pr-5px cursor-pointer border-r border-[#${
									isDark ? '181A1F' : 'DFDFDF'
								}] text-[#${
									isDark ? 'dcdcdc80' : '00000080'
								}] truncate max-w-240px hover:bg-[#${
									isDark ? '323842' : 'ffffff'
								}]`,
								state.activeTabKey === item.key
									? `bg-[#${isDark ? '2D313A' : 'FFFFFF'}] !text-[#${
											isDark ? 'dcdcdc' : '3b3b3b'
										}] ${isDark ? '' : 'border-t-1px border-t-[#005FB8]'}`
									: ``,
							]"
							:title="item.key"
							@click="handleTabChange(i)"
						>
							<div class="truncate">{{ item.label }}</div>
							<div
								:class="`w-18px h-18px rounded-[4px] flex items-center justify-center cursor-pointer ml-3 flex-shrink-0 hover:bg-[#${
									isDark ? '3F444B' : 'E9E9E9'
								}]`"
								title="关闭（Ctrl+W）"
								@click.stop="closeTab(i)"
							>
								<div class="i-ion:close"></div>
							</div>
						</div>
					</div>
					<div class="tabs-main flex-grow overflow-hidden">
						<Tab
							v-for="(item, i) in state.tabs"
							v-show="state.activeTabKey === item.key"
							:key="item.key"
							:file-path="item.key"
							:is-website-address="item.isWebsiteAddress"
							@coord-update="
								(e) => {
									if (state.activeTabKey !== item.key) {
										return
									}
									state.realTimeCoord = e
								}
							"
						/>
					</div>
				</div>
			</div>
		</div>
		<!-- 状态栏 -->
		<div
			:class="`status-bar bg-[#${
				isDark ? '21252B' : 'F8F8F8'
			}] px-3 h-20px flex items-center flex-shrink-0 border-t border-[#${
				isDark ? '282c34' : 'DFDFDF'
			}] select-none`"
		>
			<n-tooltip
				v-if="state.realTimeCoord && state.activeTabKey !== null"
				trigger="hover"
			>
				<template #trigger>
					<div
						:class="`cursor-pointer px-3 hover:bg-[#${
							isDark ? '3C3F45' : 'E9E9E9'
						}]`"
					>
						图像坐标：({{ Math.floor(state.realTimeCoord.x) }}，{{
							Math.floor(state.realTimeCoord.y)
						}})
					</div>
				</template>
				鼠标在图像中的位置
			</n-tooltip>
		</div>
		<!-- 网络地址弹窗 -->
		<n-modal
			v-model:show="state.websiteAddressModal.isModalShow"
			preset="card"
			style="width: 600px"
			title="打开网络地址"
			:mask-closable="false"
			:auto-focus="false"
		>
			<div>
				<n-input
					v-model:value="state.websiteAddressModal.websiteAddress"
					type="text"
					placeholder="例如：http://openseadragon.github.io/example-images/duomo/duomo.dzi"
					autofocus
				/>
			</div>
			<template #footer>
				<div class="flex justify-end">
					<n-button
						class="mr-3"
						type="primary"
						secondary
						@click="state.websiteAddressModal.isModalShow = false"
					>
						取消
					</n-button>
					<n-button type="primary" @click="websiteAddressSubmit()">
						确认
					</n-button>
				</div>
			</template>
		</n-modal>
	</div>
</template>

<style lang="scss" scoped>
.main {
	height: 100%;
	.resource-manager {
		.frame-size-setting {
			transition: all ease 0.25s;
			&:hover {
				cursor: e-resize;
			}
		}
	}
	.tabs {
		.tabs-bar {
		}
		.tabs-main {
		}
	}
	.status-bar {
	}
}
</style>
<style lang="scss">
.main {
	.resource-manager {
		.n-tree {
			.n-tree-node-wrapper {
				padding: 0px 0;
				.n-tree-node {
					border-radius: 0;
					align-items: center;
					.n-tree-node-indent {
						div {
							width: 14px !important;
						}
					}
					.n-tree-node-switcher {
						width: 16px !important;
						height: 16px;
						margin-left: 4px;
						margin-right: -1px;
						margin-bottom: 1px;
					}
					.n-tree-node-content {
						min-height: 22px;
						height: 22px;
						padding-left: 0;
						padding-right: 0;
						.n-tree-node-content__prefix {
							margin-right: 3px;
						}
						.n-tree-node-content__text {
							overflow: hidden;
						}
					}
				}
			}
		}
	}
}
</style>

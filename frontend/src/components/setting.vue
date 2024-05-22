<script setup lang="ts">
const props = defineProps({
	isVisible: Boolean,
})
const emits = defineEmits(['close'])
</script>
<template>
	<n-modal
		:show="props.isVisible"
		preset="card"
		style="width: 600px"
		title="设置"
		:mask-closable="false"
		:auto-focus="false"
		:close-on-esc="false"
		@close="emits('close')"
	>
		<div>
			<n-tabs type="bar" animated :placement="'left'">
				<n-tab-pane name="通用" tab="通用">
					<n-form
						ref="formRef"
						label-placement="left"
						label-width="auto"
						require-mark-placement="right-hanging"
					>
						<n-form-item label="开机自动启动" path="switchValue">
							<n-switch />
						</n-form-item>
						<n-form-item label="系统语言" path="selectValue">
							<n-select value="中文" placeholder="Select" />
						</n-form-item>
						<n-form-item label="系统主题" path="radioGroupValue">
							<n-radio-group name="radiogroup1" value="1">
								<n-space>
									<n-radio value="1"> 浅色主题 </n-radio>
									<n-radio value="2"> 深色主题 </n-radio>
								</n-space>
							</n-radio-group>
						</n-form-item>
						<n-form-item label="标签页数量" path="inputNumberValue">
							<n-input-number class="w-full" />
						</n-form-item>
						<n-form-item label="标注保存路径" path="inputNumberValue">
							<n-input />
						</n-form-item>
					</n-form>
				</n-tab-pane>
				<n-tab-pane name="外观" tab="外观"> 外观 </n-tab-pane>
				<n-tab-pane name="导出" tab="导出"> 导出 </n-tab-pane>
				<n-tab-pane name="图像" tab="图像"> 图像 </n-tab-pane>
				<n-tab-pane name="标注" tab="标注"> 标注 </n-tab-pane>
			</n-tabs>
		</div>
	</n-modal>
</template>
<style lang="scss" scoped></style>

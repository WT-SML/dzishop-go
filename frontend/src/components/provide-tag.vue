<script setup lang="ts">
import emitter from '~/tools/emitter'
import { PROVIDE_TAG, COMMON_PROVIDE } from '~/constants/evt-name'
import { TAG } from '~/constants/tag'

const state = reactive({
	provideTagPayload: {
		provideId: null,
	},
	isVisible: false,
})

// 关闭
const close = () => {
	emitter.emit(COMMON_PROVIDE, {
		provideId: state.provideTagPayload.provideId,
		result: null,
	})
	state.isVisible = false
}
// 选择
const select = (tag) => {
	emitter.emit(COMMON_PROVIDE, {
		provideId: state.provideTagPayload.provideId,
		result: tag,
	})
	state.isVisible = false
}

emitter.on(PROVIDE_TAG, (payload: any) => {
	state.provideTagPayload = payload
	state.isVisible = true
})
onUnmounted(() => {
	emitter.off(PROVIDE_TAG)
})
</script>
<template>
	<n-modal
		:show="state.isVisible"
		preset="card"
		style="width: 400px"
		title="选择标签"
		:mask-closable="false"
		:auto-focus="false"
		:close-on-esc="false"
		@close="close()"
	>
		<div>
			<n-tabs type="line" animated>
				<n-tab-pane v-for="(v, k) in TAG" :key="k" :name="k" :tab="k">
					<div class="max-h-300px overflow-y-auto">
						<div
							v-for="(item, i) in v"
							:key="i"
							class="cursor-pointer p-5px rounded hover:bg-[#005fb8] hover:text-[#fff]"
							@click="select(item)"
						>
							{{ item }}
						</div>
					</div>
				</n-tab-pane>
			</n-tabs>
		</div>
	</n-modal>
</template>
<style lang="scss" scoped></style>

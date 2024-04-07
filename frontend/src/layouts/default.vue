<script setup lang="ts">
import titleBar from '~/components/title-bar.vue'
</script>
<template>
	<div class="h-100vh flex flex-col">
		<!-- 标题栏 -->
		<titleBar class="flex-shrink-0" />
		<router-view v-slot="{ Component }" class="flex-grow">
			<transition name="fade" mode="out-in">
				<component :is="Component" />
			</transition>
		</router-view>
	</div>
</template>

<style>
.fade-enter-active,
.fade-leave-active {
	transition: opacity 0.25s ease;
}

.fade-enter-from,
.fade-leave-to {
	opacity: 0;
}
</style>

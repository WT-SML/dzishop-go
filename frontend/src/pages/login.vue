<script setup lang="ts">
import logo from '~/assets/imgs/logo.png'
</script>
<template>
	<div class="flex items-center justify-center login-container">
		<div class="shadow w-300px rounded p-8 bg-white">
			<div class="flex items-center mb-5">
				<img :src="logo" class="w-20px mr-2" />
				<div>欢迎登录 dzishop</div>
			</div>
			<n-input placeholder="请输入账号" :maxlength="32" class="mb-5">
				<template #prefix>
					<div class="i-material-symbols:person"></div>
				</template>
			</n-input>
			<n-input
				placeholder="请输入密码"
				type="password"
				:maxlength="32"
				class="mb-5"
			>
				<template #prefix>
					<div class="i-ri:lock-password-fill"></div>
				</template>
			</n-input>
			<n-button class="" type="primary" block>登录</n-button>
		</div>
	</div>
</template>

<style lang="scss" scoped>
/* 假设你有一个登录表单容器 */  
.login-container {  
  background: linear-gradient(45deg, #d9dfe4, #aeade7); 

}
</style>

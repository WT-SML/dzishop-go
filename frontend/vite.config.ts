import Tov from './presets'
import { defineConfig } from 'vite'

export default defineConfig({
	plugins: [Tov()],

	server: {
		host: '127.0.0.1',
		port: 1224,
	},
})

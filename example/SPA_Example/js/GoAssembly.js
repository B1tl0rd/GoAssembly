var ruta = "index"
/*
Copyright 2018 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

*/

		if (!WebAssembly.instantiateStreaming) { // polyfill
			WebAssembly.instantiateStreaming = async (resp, importObject) => {
				const source = await (await resp).arrayBuffer();
				return await WebAssembly.instantiate(source, importObject);
			};
		}
		const go = new Go();
		let mod, inst;
		WebAssembly.instantiateStreaming(fetch("/assembly/assembly"), go.importObject).then((result) => {
			mod = result.module;
			inst = result.instance;
			setTimeout(()=>{run()},500);
		});

		async function run() {
			await go.run(inst);
			inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
		}
		function goroute() {
			setTimeout(()=>{console.clear()},10);
			setTimeout(()=>{run()},30);
		}

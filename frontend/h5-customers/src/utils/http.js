import axios from 'axios'

const urlconfig = 'http://192.168.1.103:805'

export default {
    get(url, vm) {
        const loading = vm.$loading({
            lock: true,
            text: '玩命加载中',
            spinner: 'el-icon-loading',
            background: 'rgba(0, 0, 0, 0.7)'
        })
        return new Promise((resolve, reject) => {
            axios.get(urlconfig + url).then((res) => {
                console.log('ooooo----res start1');
                //console.log(res);

                try{
                    if (res['data']['code'] === 1) {
                        resolve(res)
                        loading.close()
                        return
                    }
                    
                    if (res['data']['code'] === 0) {
                        reject(res['data']['message'])
                        return 
                    }

                    console.log('ooooo----2');
                } catch(ex) {
                    reject(ex)
                }

                
            }).catch((rej) => {
                console.log('ooooo----3');
                reject(rej)
                loading.close()
            })
        })
    },
    post(url, params, vm) {
        console.log(vm)
        const loading = vm.$loading({
            lock: true,
            text: '玩命加载中',
            spinner: 'el-icon-loading',
            background: 'rgba(0, 0, 0, 0.7)'
        })
        return new Promise((resolve, reject) => {
            axios.post(urlconfig + url, params).then((res) => {
                console.log('ooooo----1');
                try {
                    if (res['data']['code'] === 0) {
                        resolve(res)
                        loading.close()
                    }
                    if (res['data']['code'] === 1) {
                        reject(res['data']['message'])
                        loading.close()
                    }
                    console.log('ooooo----2');
                    loading.close()
                } catch(ex) {
                    console.log('ooooo----3');
                }


            }).catch((rej) => {
                reject(rej)
                loading.close()
            })
        })
    }
}
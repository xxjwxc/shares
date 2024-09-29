<template>
    <div class="chat">
        <div v-scroll-to-bottom class="chatList scroll-hidden">
            <div class="chatListBox scroll-hidden">
                <div v-for="(item, index) in msgList" :key="index">
                    <div class="chatItem">
                        <div class="logo">
                            <el-avatar :size="30" :src="item.logo" />
                        </div>
                        <div class="msg">
                            <div class="msgTitle">
                                {{ item.title }}
                            </div>
                            <div v-if="!item.msg && loading" class="msgContent" element-loading-background="transparent" v-loading="true" style="">
                            </div>
                            <div v-else class="msgContent" :class="item.class">
                                <v-md-preview :text="item.msg"></v-md-preview>
                            </div>
                        </div>
                    </div>
                </div>
                <div v-if="!msgList.length" class="chatCard">
                    <el-avatar class="chatCardAvatar" :size="60" :src="logo" shape="square" />
                    <br />
                    <div>
                        复利Chat
                    </div>
                    <br />
                    <div class="chatCardContent">
                        你好，我是小潞，专注于投研产品化和数据科技领域，能够为资产管理和财富管理行业开发并提供专业的量化工具。
                    </div>
                </div>
            </div>
        </div>
        <div class="chatInput">
            <div class="deleteButton">
                <el-button @click="del" text type="danger" size="small">清空</el-button>
            </div>
            <input class="textarea" name="" id="" v-model="value" oninput="value=value.replace(/[\r\n]/g, '')"
                @keyup.enter="submit" />
            <span role="img" class="semi-icon semi-icon-default">
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="none" viewBox="0 0 16 16"
                    @click="submit">
                    <path fill="currentColor"
                        d="M14.464 8.903a1.027 1.027 0 0 0 0-1.805L2.508.625A1.01 1.01 0 0 0 1.5.645c-.312.186-.5.516-.5.881L2.32 6.68l5.608.82c.321 0 .581.224.581.5 0 .277-.26.5-.581.5-3.161.464-5.03.736-5.607.816L1 14.473a1.021 1.021 0 0 0 1.508.903l11.956-6.473Z">
                    </path>
                </svg>
            </span>
        </div>
        <div class="footer">Powered by <span @click="qCloud" class="isActive">QCloud.AI</span>无法确保真实准确，仅供参考。</div>
    </div>
</template>
<script setup>
import { ref } from 'vue'
import logo from '@/assets/logo.png'
const msgList = ref(JSON.parse(localStorage.getItem('msg')) || [])
const jsonData = ref([])
const value = ref('')
const loading = ref(false)
const token = 'oxFCP6nLNKUN294CLbho4QovXbxI'
const userInfo = JSON.parse(localStorage.getItem('userInfo'))

const submit = () => {
    if (loading.value) return
    msgList.value.push({
        logo: userInfo.avatarURL,
        msg: value.value,
        title: userInfo.nickName,
        class: 'user'
    })
    msgList.value.push({
        logo: logo,
        msg: '',
        title: '复利备忘录',
        class: 'app'
    })

    //定义一个EventSource对象，传入请求地址URL
    const eventSource = new EventSource(`https://www.xxxxxx.cn/shares/api/v1/analy.chat?msg=${encodeURIComponent(value.value)}&user-token=${userInfo.openid}`)
    // // 与事件源的连接刚打开时触发
    eventSource.onopen = function (e) {
        console.log(e, "连接刚打开时触发");
        value.value = ''
        loading.value = true
    };

    // 后端返回信息，格式可以和后端协商
    eventSource.addEventListener('msg', (e) => {
        const data = JSON.parse(e.data)
        msgList.value[msgList.value.length - 1].msg = msgList.value[msgList.value.length - 1].msg + data.content;
    })

    // 连接失败
    eventSource.onerror = function (e) {
        jsonData.value.forEach(i => {
            if (msgList.value[msgList.value.length - 1].msg.includes(i.value)) {
                msgList.value[msgList.value.length - 1].msg = msgList.value[msgList.value.length - 1].msg.replace(i.value, `<span class=isActive data-id="${i.key}">${i.value}</span>`)
            }
        })
        localStorage.setItem('msg', JSON.stringify(msgList.value))
        eventSource.close(); // 关闭连接
        loading.value = false
    };
}

const del = () => {
    localStorage.removeItem('msg')
    msgList.value = []
}

const qCloud = () => {
    window.parent.postMessage({
        type: 'home',
        url: 'https://www.xxxxxx.cn'
    }, '*');
}


fetch('https://www.xxxxxx.cn/shares/api/v1/file/xlsx/codes.json').then(res => res.json()).then(data => { jsonData.value = data })
document.addEventListener('click', (e) => {
    if (e.target.dataset.id) {
        // console.log(e.target.dataset.id, 'e.target.dataset.id');
        // console.log(e.target.innerText, 'e.target.innerText');
        // window.postMessage({ type:'code',value: e.target.dataset.id })
        window.parent.postMessage({
            type: 'code',
            code: e.target.dataset.id,
            name: e.target.innerText
        }, '*');
    }
})




</script>

<style>
body {
    margin: 0;
    padding: 0;
}

.isActive {
    color: #aa55ff;
    text-decoration: underline 2px solid #aa55ff;
    cursor: pointer;
}
</style>

<style scoped lang="scss">
* {
    box-sizing: border-box;
}

.chat {
    // padding: 20px;
    height: 100vh;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    position: absolute;
    top: 0;
    left: 0;

    .chatCard {
        text-align: center;
        position: absolute;
        top: 50%;
        left: 50%;
        color: #aa55ff;
        transform: translate(-50%, -50%);
        &Avatar {
            background-color: transparent;
            border: 1px solid #aa55ff;
            border-radius: 12px;
        }

        &Content {
            background-color: rgb(60, 60, 60);
            width: 300px;
            margin: auto;
            text-align: left;
            padding: 10px;
            color: darkgrey;
            border-radius: 12px;
            border: 1px solid rgb(60, 60, 60);
        }
    }

    .chatList {
        width: 100%;
        flex: 1;
        overflow: auto;

        .chatListBox {
            margin: 0px auto;
            width: 100vw;

            .chatItem {
                display: flex;
                margin-bottom: 30px;

                .logo {
                    width: 36px;
                }

                .msg {
                    flex: 1;

                    .msgTitle {
                        font-weight: 600;
                        font-size: 14px;
                        color: darkgrey;
                    }

                    .msgContent {
                        margin-top: 5px;
                        display: flex;
                        width: -moz-fit-content;
                        width: fit-content;
                        padding: 10px;
                        border-radius: 10px;
                        // border: 1px solid #ddd;
                    }
                }
            }

            :deep(.vuepress-markdown-body tr:nth-child(2n)) {
                background-color: inherit;


            }

            :deep(.vuepress-markdown-body) {

                th,
                td,
                tr {
                    // border: none;
                    border-color: gray;
                }

                &:not(.custom) {
                    // padding: 0px;
                    // background-color: transparent;
                    // // color: #745605;
                    // color: slategrey; // inherit;
                    // font-size: 15px;


                    padding: 0px;
                    background-color: inherit;
                    color: inherit;
                    // font-size:14px;
                    font-size: 0.8rem;
                }
            }
        }

    }

    .user {
        background-color: darkgreen;
        color: #fff !important;
    }

    .app {
        background-color: darkgray; // seashell,darkgray
        color: #000;
    }


    .chatInput {
        position: relative;
        width: 100vw;
        margin: 0px auto;
        display: flex;
        align-items: center;

        .deleteButton {

            // width: 40px;
            :deep(.el-button) {
                color: #c7c9cc;

                &:hover {
                    background-color: transparent;
                }
            }
        }

        .semi-icon {
            position: absolute;
            right: 10px;
            top: 50%;
            transform: translate(0, -45%);
            fill: currentColor;
            display: inline-block;
            font-style: normal;
            text-align: center;
            text-rendering: optimizeLegibility;
            text-transform: none;
            font-size: 16px;
            color: #aa55ff;
        }

        .textarea {
            width: 100%;
            padding: 12px;
            color: rgb(204, 204, 172);
            border: 1px solid rgb(60, 60, 60);
            border-radius: 8px;
            padding-right: 30px;
            background-color: rgb(60, 60, 60);

            &:focus,
            &:focus-within,
            &:active,
            &:visited {
                border: 1px solid rgb(60, 60, 60) !important;
            }

            &:hover,
            &:focus-visible {
                border: 1px solid #aa55ff !important;
                outline-offset: 0px;
                outline: 1px solid #aa55ff;
            }
        }
    }

    .footer {
        color: rgba(204, 204, 172, .7);
        font-size: 12px;
        font-weight: 400;
        line-height: 16px;
        padding: 8px 0px;
        text-align: center;
        width: 100%;
    }


    /* 在你的全局样式文件中添加 */
    .scroll-hidden {
        overflow: -moz-scrollbars-none;
        /* Firefox */
        -ms-overflow-style: none;
        /* IE 10+ */
        scrollbar-width: none;
        /* Firefox */
        overflow-y: scroll;
    }

    .scroll-hidden::-webkit-scrollbar {
        /* WebKit browsers */
        display: none;
    }
}
</style>
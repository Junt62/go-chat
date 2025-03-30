<template>
  <el-aside width="240px">
    <div class="content">
      <div class="title">用户设置</div>
      <MessageButton title="我的账号" path="/chat/online" />
      <MessageButton title="个人资料" path="/chat/friend" />
      <MessageButton title="内容与社交" path="/chat/friend" />
      <MessageButton title="数据与隐私" path="/chat/friend" />
      <MessageButton title="家庭中心" path="/chat/friend" />
      <MessageButton title="已授权的APP" path="/chat/friend" />
      <MessageButton title="设备" path="/chat/friend" />
      <MessageButton title="连接" path="/chat/friend" />
      <MessageButton title="剪辑片段" path="/chat/friend" />

      <el-divider title="divider" />

      <div class="title">账单设置</div>
      <MessageButton title="Nitro" path="/chat/friend" />
      <MessageButton title="服务器助力" path="/chat/zhangsan" />
      <MessageButton title="订阅" path="/chat/lisi" />
      <MessageButton title="礼物库" path="/chat/wangwu" />
      <MessageButton title="账单" path="/chat/wangwu" />

      <el-divider title="divider" />

      <div class="title">APP设置</div>
      <MessageButton title="外观" path="/chat/wangwu" />
      <MessageButton title="可访问性" path="/chat/wangwu" />
      <MessageButton title="语音和视频" path="/chat/wangwu" />
      <MessageButton title="聊天" path="/chat/wangwu" />
      <MessageButton title="通知" path="/chat/wangwu" />
      <MessageButton title="快捷键" path="/chat/wangwu" />
      <MessageButton title="语言" path="/chat/wangwu" />
      <MessageButton title="主播模式" path="/chat/wangwu" />
      <MessageButton title="高级设置" path="/chat/wangwu" />

      <el-divider title="divider" />

      <div class="title">活动设置</div>
      <MessageButton title="当前状态隐私" path="/chat/wangwu" />

      <el-divider title="divider" />

      <MessageButton title="新鲜事儿" path="/chat/wangwu" />
      <MessageButton title="周边商品" path="/chat/wangwu" />
      <MessageButton title="HypeSquad" path="/chat/wangwu" />

      <el-divider title="divider" />

      <MessageButton title="登出" path="/chat/wangwu" />

      <p><br /></p>
    </div>
  </el-aside>
</template>

<script setup lang="ts">
import MessageButton from '@/components/MsgBtn.vue'
</script>

<style scoped>
.content {
  height: 100vh;
  padding: 10px;
  overflow-y: auto;
  background-color: rgba(0, 0, 0, 0.2);
}
.title {
  margin-bottom: 12px;
  color: rgba(255, 255, 255, 0.8);
}
</style>

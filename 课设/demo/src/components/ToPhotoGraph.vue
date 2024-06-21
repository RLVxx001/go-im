<template>
  <div style="">
    <hr>
    <div style="display:flex;width:800px;height:30px;margin-bottom:-10px;margin-top:-10px;margin-left:70px">
      <router-link to="/toSpace" style="text-decoration: none;">
        <div class="to1" style="margin-top:-7px">发表</div>
      </router-link>
      <router-link to="/toMessage" style="text-decoration: none;">
        <div class="to1" style="margin-top:-7px">留言</div>
      </router-link>
      <router-link to="" style="text-decoration: none;">
        <div class="to1" style="margin-top:-7px">与我相关</div>
      </router-link>
      <router-link to="/toPhotoGraph" style="text-decoration: none;">
        <div class="to1" style="margin-top:-7px">相册</div>
      </router-link>
      <router-link to="" style="text-decoration: none;">
        <div class="to1" style="margin-top:-7px">日志</div>
      </router-link>
    </div>
    <hr>
  <el-scrollbar style="width:800px;height:470px;margin-top:-20px" ref="scrollbarRef" always>
      <div style="margin-top:20px;">
      <el-form-item label="精选照片">
      <el-upload 
            class="dl-avatar-uploader-min square" 
            :class="{uoloadBtn:showBtnDealImg,disUoloadBtn:noneBtnImg}"
            :http-request="httpRequest"
            :limit="limitCountImg"
            :on-remove="coverFileRemove"
            :on-exceed="handleExceedCover"
            :before-upload="beforeImageUpload"
            :on-change="handleImgChange"
            :on-preview="handlePictureCardPreview"
            v-model:file-list="fileList"
            list-type="picture-card"
            accept="image/*"
            multiple
            >
            <el-button type="text" style="font-size:50px"></el-button>
            <template #tip>
                <div class="el-upload__tip" style="color:rgba(220, 228, 253, 0.942);word-break:break-all;width:120px">最多上传100张图片,且单张图片大小不能超过10MB</div>
            </template>
        </el-upload>
        <el-dialog v-model="dialogVisible">
          <img w-full :src="dialogImageUrl" alt="Preview Image" style="width: 100%;" />
        </el-dialog>
      </el-form-item>
    </div>
    </el-scrollbar>
  </div>
</template>
<script setup>
import { ref, onMounted ,h,reactive,nextTick, isRef } from 'vue'; 
import { ElNotification,ElScrollbar } from 'element-plus'
import service from '../axios-instance'
import axios from "axios";
const $MYGO = inject('$MYGO', '');
let rcd=reactive([])
let id = ref(localStorage.getItem("id"));
onMounted(()=>{
  getimglist()
})
function Detail(e){
  localStorage.setItem("ToId",e-0)

}

let fileList=reactive([])
let userImgList=reactive([])
let	limitCountImg=ref(100)
let	showBtnDealImg=ref(true)
let noneBtnImg=ref(false)
const dialogImageUrl = ref('')
const dialogVisible = ref(false)
const dataSet = reactive({})
const uploadUrl=ref($MYGO+'/userImg/upload')
const addform=reactive({image:''})
function httpRequest(option){
  let dataForm = new FormData();
  dataForm.append('file',option.file)
  dataForm.append('uid',option.file.uid)
  axios({
        method: 'POST',
        url: uploadUrl.value,
        data: dataForm,
//设置请求参数的规则
        headers: {
            "Content-Type": "multipart/form-data",
            "Authorization":localStorage.getItem('token')
        }
    }).then(response => {

        userImgList.push(response.data)
        userImgList[userImgList.length-1].uid=option.file.uid
    }).catch(err=>{
      console.error(err)
    })

}
function coverFileRemove(file, fileList) {

    if(file.id)
    {
      for(let i=0;i<userImgList.length;i++)
      {
        if(file.id==userImgList[i].id)
        {
          service.post($MYGO+'/userImg/delete',{'id':userImgList[i].id-0})
          .then(response => {
            userImgList.splice(i,1)
          }).catch(err=>{
            console.error(err)
            })
          return
        }
      }
    }
    else
    {
      for(let i=userImgList.length-1;i>=0;i--)
      {
        if(file.uid==userImgList[i].uid)
        {
          service.post($MYGO+'/userImg/delete',{id:userImgList[i].id-0})
          .then(response => {
            userImgList.splice(i,1)
          }).catch(err=>{
            console.error(err)
            })
          return
        }
      }
    }
}
function handlePictureCardPreview(uploadFile){
  dialogImageUrl.value = uploadFile.url
  dialogVisible.value = true
}
function beforeImageUpload(rawFile){
    if(rawFile.size / 1024 / 1024 > 10){
        ElMessage.error("单张图片大小不能超过10MB!");
        return false;
    }
    return true;
}
function handleExceedCover(files, fileList){
    ElMessage.error({
        message: `上传图片数量超出限制！`,
        type: "error",
    });
}

function handleImgChange(file, fileList){
    noneBtnImg.value = fileList.length >= limitCountImg.value;
}
function getimglist(){
  service.post($MYGO+'/userImg/getByFriend',{
    "userId":localStorage.getItem("toId")-0,
    "id":localStorage.getItem("id")-0,
    "url":"#"
  })
  .then(response => {
      if(!response.data){
        return
      }
      response.data.forEach(element => {
        fileList.push(element)
        userImgList.push(element)
      });
    }).catch(err=>{
      console.error(err)
    })
}
</script>
<style>
.to22{
  color:rgb(111, 111, 111);
}
</style>
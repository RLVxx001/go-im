<template>
  <div>
    <ul style="color:rgb(207, 234, 244);margin-left:150px;margin-top:80px;font-size:25px">
      <div style="display:flex;">
        <img :src="user.img" style="height:100px;width:100px;border-radius:50%"  />
        <input type="file" style="margin-left:40px;margin-top:60px;font-size:15px" />
      </div>
      <li style="margin-top:20px">账号：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.username"/></li>
      <li style="margin-top:10px">用户名：<input style="margin-left:26px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.account"/></li>
      <li style="margin-top:10px">个性签名：<input style="background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.signed"/></li>
      <li style="margin-top:10px">邮箱：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.email"/></li>
      <li style="margin-top:10px">生日：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.birthday"/></li>
      <li style="margin-top:10px">居住地：<input style="margin-left:24px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.city"/></li>
    </ul>
    <div>
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
            <el-button type="text" style="font-size:50px">+</el-button>
            <template #tip>
                <div class="el-upload__tip">最多上传100张图片,且单张图片大小不能超过10MB</div>
            </template>
        </el-upload>
        <el-dialog v-model="dialogVisible">
          <img w-full :src="dialogImageUrl" alt="Preview Image" style="width: 100%;" />
        </el-dialog>
      </el-form-item>
    </div>
    <div>
      <button @click="save()" style="margin-top:20px;margin-left:300px;width:110px;height:40px;font-size:20px;line-height:40px;border-radius:10px">保存</button>
    </div>
  </div>
</template>
<script setup>
import {ref,reactive,onMounted} from 'vue'
import { ElNotification } from 'element-plus'
import service from '../axios-instance'
import axios from "axios";
import { useWsStore } from '../store/user';
const wsStore=useWsStore()
// const user=reactive(JSON.parse(localStorage.getItem('user')))
var user = reactive(JSON.parse(localStorage.getItem('user')))
onMounted(()=>{
  wsStore.event=-1
  getimglist()
  // service.post("http://localhost:8080/user/fidUser",{"username":localStorage.getItem("username")})
  // .then(tmp=>{
  //   user.username=tmp.data.username
  //   user.account=tmp.data.account
  //   user.signed=tmp.data.signed
  //   user.email=tmp.data.email
  //   user.pht=tmp.data.img
  //   user.birthday=tmp.data.birthday
  //   user.city=tmp.data.city
  // })
})
function save()
{
  console.log(user)
  service.post("http://localhost:8080/user/update",{
    'account':user.account,
    'signed':user.signed,
    'birthday':user.birthday
  })
  .then(res=>{
    localStorage.setItem('user',JSON.stringify(user))
    ElNotification({
      title: 'Success',
      message: '修改成功',
      type: 'success',
    }) 
  }).catch(err=>{
    ElNotification({
      title: 'Error',
      message: '修改失败',
      type: 'error',
    })
  })
}

let fileList=reactive([])
let userImgList=reactive([])
let	limitCountImg=ref(100)
let	showBtnDealImg=ref(true)
let noneBtnImg=ref(false)
const dialogImageUrl = ref('')
const dialogVisible = ref(false)
const dataSet = reactive({})
const uploadUrl=ref('http://localhost:8080/userImg/upload')
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
          service.post('http://localhost:8080/userImg/delete',{'id':userImgList[i].id-0})
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
          service.post('http://localhost:8080/userImg/delete',{id:userImgList[i].id-0})
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
  service.get('http://localhost:8080/userImg/getByUser')
  .then(response => {
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
.disUoloadBtn .el-upload--picture-card{
    display:none;   /* 上传按钮隐藏 */
}
</style>
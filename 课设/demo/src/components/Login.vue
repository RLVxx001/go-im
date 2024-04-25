<template>
  <div style="background-color: #525257a2;width:1500px;height:800px">
    <div style="width:400px;height:400px;position:absolute;left:550px;top:150px;">
      <form>
        <div class="form-group">
          <label for="exampleInputEmail1" style="color:rgba(220, 228, 253, 0.942);">Email address</label>
          <input type="email" class="form-control"  id="exampleInputEmail1" placeholder="Email">
        </div>
        <div class="form-group">
          <label for="exampleInputPassword1">Password</label>
          <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
        </div>
        <div class="form-group">
          <label for="exampleInputFile">File input</label>
          <input type="file" id="exampleInputFile">
          <p class="help-block">Example block-level help text here.</p>
        </div>
        <div class="checkbox">
          <label>
            <input type="checkbox"> Check me out
          </label>
        </div>
        <button type="submit" class="btn btn-default">Submit</button>
      </form>
    </div>
  </div>
</template>

<script>

  export default{
  components:{

  },
}
</script>


<style>

</style>
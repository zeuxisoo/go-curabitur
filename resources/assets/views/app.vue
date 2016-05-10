<template>
    <div id="app">
        <div class="container">
            <div class="navbar navbar-default" role="navigation">
                <div class="container-fluid">
                    <div class="navbar-header">
                        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
                            <span class="sr-only">Toggle navigation</span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                        </button>
                        <a class="navbar-brand">Brand</a>
                    </div>
                    <div class="navbar-collapse collapse">
                        <ul class="nav navbar-nav">
                            <li>
                                <a href="">
                                    <i class="glyphicon glyphicon-home"></i> Home
                                </a>
                            </li>
                        </ul>
                        <ul class="nav navbar-nav navbar-right">
                            <li>
                                <a class="nav-link">
                                    <i class="glyphicon glyphicon-scale"></i> About
                                </a>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
            <router-view></router-view>
        </div>
    </div>
</template>

<style>
html, body {
    height: 100%;
}

#app {
    height: 100%;
}

#app .container {
    height: 100%;
}
</style>

<script>
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min'

import 'sweetalert/dist/sweetalert.css'
import 'sweetalert'
</script>

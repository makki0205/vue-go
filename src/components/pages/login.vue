<template>
    <form action="">
        <div class="modal-card" style="width: auto">
            <header class="modal-card-head">
                <p class="modal-card-title">ログイン</p>
            </header>
            <section class="modal-card-body">
                <b-field label="Email">
                    <b-input
                            type="email"
                            v-model="email"
                            placeholder="Your email"
                            required>
                    </b-input>
                </b-field>

                <b-field label="Password">
                    <b-input
                            type="password"
                            v-model="password"
                            password-reveal
                            placeholder="Your password"
                            required>
                    </b-input>
                </b-field>

                <b-checkbox>Remember me</b-checkbox>
            </section>
            <footer class="modal-card-foot">
                <button class="button" type="button">新規登録</button>
                <button class="button is-primary" @click="login">ログイン</button>
            </footer>
        </div>
    </form>
</template>

<script>
    import auth from"../../service/auth"
    export default {
        data() {
            return {
                email: "",
                password: ""
            }
        },
        methods:{
            login(){

            }
        }
    }
</script>
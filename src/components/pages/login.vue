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
                            :value="email"
                            placeholder="Your email"
                            required>
                    </b-input>
                </b-field>

                <b-field label="Password">
                    <b-input
                            type="password"
                            :value="password"
                            password-reveal
                            placeholder="Your password"
                            required>
                    </b-input>
                </b-field>

                <b-checkbox>Remember me</b-checkbox>
            </section>
            <footer class="modal-card-foot">
                <button class="button" type="button">新規登録</button>
                <button class="button is-primary">ログイン</button>
            </footer>
        </div>
    </form>
</template>

<script>
    export default {
    }
</script>
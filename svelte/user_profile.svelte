<script>
    // @ts-nocheck
    import Menu from './_components/Menu.svelte';
    import ProfileHeader from './_components/ProfileHeader.svelte';
    import Footer from './_components/Footer.svelte';
    import {notifier} from './_components/notifier.js'
    import Icon from 'svelte-icons-pack/Icon.svelte';
    import {datetime} from './_components/formatter';
    import {onMount} from 'svelte';
    import {UserChangePassword, UserSessionKill, UserSessionsActive, UserUpdateProfile} from './jsApi.GEN.js';
    import FaSolidAngleRight from 'svelte-icons-pack/fa/FaSolidAngleRight';
    import FaSolidTimes from 'svelte-icons-pack/fa/FaSolidTimes';
    import FaSolidCircleNotch from "svelte-icons-pack/fa/FaSolidCircleNotch";

    let user = {/* user */};
    let segments = {/* segments */};
    let sessionActiveLists = {/* activeSessions */};
    let oldPassword = '';
    let newPassword = '';
    let repeatNewPassword = '';
    let oldProfileJson = '';
    let profileSubmit = false, passwordSubmit = false;
    onMount(async () => {
        console.log('onMount.user')
        oldProfileJson = JSON.stringify(user);
        console.log('User data = ', user)
    });

    async function updateProfile() {
        profileSubmit = true
        if (JSON.stringify(user) === oldProfileJson) return notifier.showWarning('No changes');
        await UserUpdateProfile(user, function(res) {
            profileSubmit = false;
            if (res.error) return notifier.showError(res.error);
            oldProfileJson = JSON.stringify(res.user);
            user = res.user;
            notifier.showSuccess('Profile updated');
        });
    }

    async function changePassword() {
        passwordSubmit = true;
        if (newPassword !== repeatNewPassword) {
            passwordSubmit = false;
            return notifier.showWarning('New password and repeat new password must be same');
        }
        let input = {
            oldPass: oldPassword,
            newPass: newPassword,
        };
        await UserChangePassword(input, function(res) {
            passwordSubmit = false;
            if (res.error) return notifier.showError(res.error);
            oldPassword = '';
            newPassword = '';
            repeatNewPassword = '';
            notifier.showSuccess('Password updated');
        });
    }

    async function killSession(sessionToken) {
        await UserSessionKill({sessionTokenHash: sessionToken}, async res => {
            if (res.error) return notifier.showError(res.error);
            if (res.sessionTerminated < 1) return notifier.showError('No session terminated');
            notifier.showInfo(res.sessionTerminated + ' session terminated');
            await UserSessionsActive({userId: user.id}, res => {
                if (res.error) return notifier.showError(res.error);
                sessionActiveLists = res.sessionsActive;
            });
        });
    }
</script>


<section class='dashboard'>
    <Menu access={segments}/>
    <div class='dashboard_main_content'>
        <ProfileHeader {user} access={segments}/>
        <div class='content'>
            <div class='profile_details_container'>
                <div class='left'>
                    <div class='profile_details'>
                        <h2>Profile Details</h2>
                        <!--						<div class="profile_pictures">-->
                        <!--							<div class="img_container">-->
                        <!--								<img alt="profile" src="/assets/img/team-1-200x200.jpg"/>-->
                        <!--							</div>-->
                        <!--							<div class="actions">-->
                        <!--								<button class='btn_upload_photo'>Upload Profile Photo</button>-->
                        <!--								<button class='btn_delete'>-->
                        <!--									<Icon color="#FFF" size={15} src={FaSolidTrashAlt}/>-->
                        <!--								</button>-->
                        <!--							</div>-->
                        <!--						</div>-->
                        <div class='input_container'>
                            <div class='name'>
                                <div class='profile_input'>
                                    <label for='fullName'>Full Name</label>
                                    <input bind:value={user.fullName} id='fullName' type='text'/>
                                </div>
                                <div class='profile_input email'>
                                    <label for='email'>E-Mail</label>
                                    <input bind:value={user.email} id='email' type='email'/>
                                </div>
                            </div>


                        </div>
                        <div class='info_container'>
                            <div class='profile_info'>
                                <label for='registered'>Registered:</label>
                                <span id='registered'>{datetime(user.createdAt)}</span>
                            </div>
                            <div class='profile_info'>
                                <label for='lastLogin'>Last login:</label>
                                <span id='lastLogin'>{datetime(user.lastLoginAt)}</span>
                            </div>
                            <div class='profile_info'>
                                <label for='verified'>Verified:</label>
                                <span id='verified'>{datetime(user.verifiedAt) || '0'}</span>
                            </div>
                        </div>
                        <label for='updateProfile'/>
                        <button id='updateProfile' on:click={updateProfile}>
                            <span>SUBMIT</span>
                            {#if !profileSubmit}
                                <Icon color='#FFF' size={18} src={FaSolidAngleRight}/>
                            {:else}
                                <Icon className="spin" color='#FFF' size={18} src={FaSolidCircleNotch}/>
                            {/if}
                        </button>
                    </div>
                    <div class='session_list_container'>
                        <h2>Active Sessions</h2>
                        <div class='session_list_header'>
                            <span>IP Address</span>
                            <span>Expired At</span>
                            <span>Device</span>
                        </div>
                        <div class='session_list'>
                            {#if sessionActiveLists.length}
                                {#each sessionActiveLists as session}
                                    <div class='session'>
                                        <span>{session.loginIPs || 'no-data'}</span>
                                        <span>{datetime(session.expiredAt) || 0}</span>
                                        <span>{session.device || 'no-data'}</span>
                                        <button on:click={() => killSession(session.sessionToken)} class='kill_session' title='Kill this session'>
                                            <Icon color='#FFF' size={12} src={FaSolidTimes}/>
                                        </button>
                                    </div>
                                {/each}
                            {/if}
                        </div>
                    </div>
                </div>

                <div class='right'>
                    <div class='password_set'>
                        <h2>Change Password</h2>
                        <div class='input_container'>
                            <div class='profile_input'>
                                <label for='oldPassword'>Old Password</label>
                                <input bind:value={oldPassword} id='oldPassword' type='password'/>
                            </div>
                            <div class='profile_input'>
                                <label for='newPassword'>New Password</label>
                                <input bind:value={newPassword} id='newPassword' type='password'/>
                            </div>
                            <div class='profile_input'>
                                <label for='repeatNewPassword'>Repeat New Password</label>
                                <input bind:value={repeatNewPassword} id='repeatNewPassword' type='password'/>
                            </div>
                            <label for='changePassword'/>
                            <button id='changePassword' on:click={changePassword}>
                                <span>SUBMIT</span>
                                {#if !passwordSubmit}
                                    <Icon color='#FFF' size={18} src={FaSolidAngleRight}/>
                                {/if}
                                {#if passwordSubmit}
                                    <Icon className="spin" color='#FFF' size={18} src={FaSolidCircleNotch}/>
                                {/if}
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <Footer/>
    </div>
</section>

<style>
    @keyframes spin {
        from {
            transform : rotate(0deg);
        }
        to {
            transform : rotate(360deg);
        }
    }

    :global(.spin) {
        animation : spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
    }

    .profile_details_container {
        position              : relative;
        margin-top            : -40px;
        margin-left           : auto;
        margin-right          : auto;
        display               : grid;
        grid-template-columns : 70% auto;
        gap                   : 30px;
        width                 : 88%;
        color                 : #475569;
        font-size             : 14px;
    }

    .profile_details_container h2 {
        font-size   : 18px;
        margin      : 0 0 20px 0;
        font-weight : 700;
        text-align  : center;
        color       : #6366F1;
    }

    .profile_details_container .left .profile_details,
    .profile_details_container .right .password_set,
    .profile_details_container .left .session_list_container {
        display          : flex;
        flex-direction   : column;
        border-radius    : 8px;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        padding          : 20px;
        background-color : white;
        height           : fit-content;
    }

    .profile_details_container .left,
    .profile_details_container .right {
        display        : flex;
        flex-direction : column;
        gap            : 30px;
    }

    .profile_details_container .left .profile_details .input_container,
    .profile_details_container .right .password_set .input_container {
        display        : flex;
        flex-direction : column;
        gap            : 10px;
    }

    .profile_details_container .left .profile_details .input_container .name {
        display               : grid;
        grid-template-columns : 1fr 1fr;
        gap                   : 20px;
    }

    .profile_details_container .left .profile_details .name .profile_input {
        width : 100%;
    }

    .profile_details_container .left .profile_details .profile_input.email {
        width        : 100%;
        margin-right : 10px;
    }

    .profile_details_container .info_container {
        display        : flex;
        flex-direction : column;
        gap            : 5px;
        margin-top     : 15px;
    }

    .profile_details_container .info_container .profile_info {
        display     : inline-flex;
        align-items : center;
        font-size   : 13px;
        margin-left : 10px;
    }

    .profile_details_container .info_container .profile_info label {
        font-weight  : 700;
        margin-right : 10px;
    }

    .profile_details_container .left .session_list_container .session_list_header {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        font-weight     : bold;
        padding         : 15px 0;
        border-bottom   : 1px solid #CBD5E1;
        margin-right    : 30px;
    }

    .profile_details_container .left .session_list_container .session_list {
        display        : flex;
        flex-direction : column;
        gap            : 5px;
        margin-right   : 30px;
    }

    .profile_details_container .left .session_list_container .session_list .session {
        text-align      : left;
        display         : flex;
        flex-direction  : row;
        align-items     : center;
        justify-content : space-between;
        padding         : 15px 0;
        position        : relative;
    }

    .profile_details_container .left .session_list_container .session_list .session span:nth-child(3),
    .profile_details_container .left .session_list_container .session_list_header span:nth-child(3) {
        flex-grow : 1;
    }

    .profile_details_container .left .session_list_container .session_list .session span,
    .profile_details_container .left .session_list_container .session_list_header span {
        width : 200px;
    }

    .profile_details_container .left .session_list_container .session_list .session .kill_session {
        border           : none;
        background-color : #EF4444;
        padding          : 6px;
        border-radius    : 50%;
        position         : absolute;
        right            : -30px;
        cursor           : pointer;
    }

    .profile_details_container .left .session_list_container .session_list .session .kill_session:hover {
        background-color : #F85454;
    }


    .profile_details #updateProfile,
    .password_set #changePassword {
        margin-left      : auto;
        width            : fit-content;
        filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        padding          : 6px 20px;
        font-size        : 13px;
        display          : inline-flex;
        font-weight      : 600;
        flex-direction   : row;
        align-items      : center;
        align-content    : center;
        justify-content  : center;
        border           : none;
        background-color : #6366F1;
        border-radius    : 8px;
        color            : white;
        cursor           : pointer;
        gap              : 6px;
    }

    .profile_details #updateProfile:hover,
    .password_set #changePassword:hover {
        background-color : #7E80F1;
    }

    /* Input box */
    .profile_input {
        display        : flex;
        flex-direction : column;
        gap            : 8px;
        width          : 100%;
    }

    .profile_input label {
        font-size   : 13px;
        font-weight : 700;
        margin-left : 10px;
    }

    /* Responsive to mobile device */
    @media (max-width : 768px) {
        .profile_details_container {
            display        : flex;
            flex-direction : column-reverse;
            gap            : 20px;
        }

        .profile_details_container .left {
            gap : 20px;
        }

        .profile_details_container .left .profile_details .input_container .name {
            display        : flex;
            flex-direction : column;
            gap            : 20px;
        }
    }
</style>
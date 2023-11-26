<script>
  // @ts-nocheck
  import SideMenu from './_components/partials/SideMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import Footer from './_components/partials/Footer.svelte';
  import { notifier } from './_components/notifier.js';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import { datetime } from './_components/formatter';
  import { onMount } from 'svelte';
  import { UserChangePassword, UserSessionKill, UserSessionsActive, UserUpdateProfile } from './jsApi.GEN.js';
  import FaSolidTimes from 'svelte-icons-pack/fa/FaSolidTimes';
  import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';
  import { localeDatetime } from './_components/formatter.js';

  let user = {/* user */};
  let segments = {/* segments */};
  let sessionActiveLists = {/* activeSessions */};
  let oldPassword = '';
  let newPassword = '';
  let repeatNewPassword = '';
  let oldProfileJson = '';
  let profileSubmit = false,
    passwordSubmit = false;
  onMount(async () => {
    console.log('onMount.user');
    oldProfileJson = JSON.stringify(user);
    console.log('User data = ', user);
  });

  async function updateProfile() {
    profileSubmit = true;
    if (JSON.stringify(user) === oldProfileJson) return notifier.showWarning('No changes');
    await UserUpdateProfile(user, function (res) {
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
    await UserChangePassword(input, function (res) {
      passwordSubmit = false;
      if (res.error) return notifier.showError(res.error);
      oldPassword = '';
      newPassword = '';
      repeatNewPassword = '';
      notifier.showSuccess('Password updated');
    });
  }

  async function killSession(sessionToken) {
    await UserSessionKill({ sessionTokenHash: sessionToken }, async res => {
      if (res.error) return notifier.showError(res.error);
      if (res.sessionTerminated < 1) return notifier.showError('No session terminated');
      notifier.showInfo(res.sessionTerminated + ' session terminated');
      await UserSessionsActive({ userId: user.id }, res => {
        if (res.error) return notifier.showError(res.error);
        sessionActiveLists = res.sessionsActive;
      });
    });
  }
</script>

<div class="root_layout">
  <SideMenu access={segments} />
  <div class="root_container">
    <Navbar {user} />
    <div class="root_content">
      <div class="user_details_container">
        <div class="profile_details">
          <h3>Profile Details</h3>
          <div class="input_row">
            <InputBox id="fullname" label="Full Name" value={user.fullName} type="text" />
            <InputBox id="email" label="E-Mail" value={user.email} type="email" />
          </div>
          <div class="user_info">
            <div class="info">
              <span>Join At</span>
              <span>{localeDatetime(user.createdAt)}</span>
            </div>
            <div class="info">
              <span>User ID</span>
              <span>{user.id}</span>
            </div>
            <div class="info">
              <span>Role</span>
              <span>{user.role}</span>
            </div>
            <div class="info">
              <span>Last Login</span>
              <span>{localeDatetime(user.lastLoginAt)}</span>
            </div>
            <div class="info">
              <span>Updated At</span>
              <span>{localeDatetime(user.updatedAt) || 0}</span>
            </div>
            <div class="info">
              <span>Verified At</span>
              <span>{localeDatetime(user.verifiedAt) || 0}</span>
            </div>
            <div class="info">
              <span>Last Updated Password</span>
              <span>{localeDatetime(user.passwordSetAt) || 0}</span>
            </div>
          </div>
          <SubmitButton on:click={updateProfile} isSubmitted={profileSubmit} />
        </div>
        <div class="password_set">
          <h3>Change Password</h3>
          <InputBox id="oldPassword" label="Old Password" value={oldPassword} type="password" />
          <InputBox id="newPassword" label="New Password" value={newPassword} type="password" />
          <InputBox id="repeatNewPassword" label="Repeat New Password" value={repeatNewPassword} type="password" />
          <SubmitButton on:click={changePassword} isSubmitted={passwordSubmit} />
        </div>
      </div>

      <div class="session_list_container">
        <h3>Active Sessions</h3>

        <div class="session_list_header">
          <span>IP Address</span>
          <span>Expired At</span>
          <span>Device</span>
        </div>

        <div class="session_list">
          {#if sessionActiveLists.length}
            {#each sessionActiveLists as session}
              <div class="session">
                <span>{session.loginIPs || 'no-data'}</span>
                <span>{datetime(session.expiredAt) || 0}</span>
                <span>{session.device || 'no-data'}</span>
                <button
                  on:click={() => killSession(session.sessionToken)}
                  class="kill_session"
                  title="Kill this session"
                >
                  <Icon color="#FFF" size={12} src={FaSolidTimes} />
                </button>
              </div>
            {/each}
          {/if}
        </div>
      </div>
    </div>

    <Footer />
  </div>
</div>

<style>
  /* +============= Global ==============+ */
  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  .user_details_container h3 {
    margin: 0;
    font-size: var(--font-md);
    width: 100%;
    text-align: center;
  }

  /* +============ /// ===========+ */

  .user_details_container {
    display: flex;
    flex-direction: row;
    gap: 30px;
  }

  .user_details_container .profile_details,
  .user_details_container .password_set {
    padding: 20px;
    border: 1px solid var(--gray-003);
    display: flex;
    flex-direction: column;
    gap: 20px;
    height: fit-content;
    border-radius: 8px;
  }

  .user_details_container .profile_details {
    width: 700px;
  }

  .user_details_container .profile_details .input_row {
    display: flex;
    flex-direction: row;
    gap: 20px;
  }

  .user_details_container .profile_details .user_info {
    border: 1px solid var(--gray-003);
    display: flex;
    flex-direction: column;
    height: fit-content;
    border-radius: 8px;
  }

  .user_details_container .profile_details .user_info .info {
    border-bottom: 1px solid var(--gray-003);
    padding: 10px 15px;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    text-transform: capitalize;
  }

  .user_details_container .profile_details .user_info .info span:first-child {
    font-weight: 700;
    width: 200px;
  }

  .user_details_container .profile_details .user_info .info:nth-child(odd) {
    background-color: var(--gray-001);
  }

  .user_details_container .profile_details .user_info .info:last-child {
    border-bottom: none;
  }

  .user_details_container .password_set {
    width: 400px;
  }

  /* +=======================================+ */

  .session_list_container .session_list_header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    font-weight: bold;
    padding: 15px 0;
    border-bottom: 1px solid #cbd5e1;
    margin-right: 30px;
  }

  .session_list_container .session_list {
    display: flex;
    flex-direction: column;
    gap: 5px;
    margin-right: 30px;
  }

  .session_list_container .session_list .session {
    text-align: left;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 15px 0;
    position: relative;
  }

  .session_list_container .session_list .session span:nth-child(3),
  .session_list_container .session_list_header span:nth-child(3) {
    flex-grow: 1;
  }

  .session_list_container .session_list .session span,
  .session_list_container .session_list_header span {
    width: 200px;
  }

  .session_list_container .session_list .session .kill_session {
    border: none;
    background-color: #4444ef;
    padding: 6px;
    border-radius: 50%;
    position: absolute;
    right: -30px;
    cursor: pointer;
  }

  .session_list_container .session_list .session .kill_session:hover {
    background-color: #f85454;
  }

  /* Responsive to mobile device */
  @media (max-width: 768px) {
  }
</style>

<script>
  /** @typedef {import('./_components/types/user').User} User */
  /** @typedef {import('./_components/types/access').Access} Access */
  /** @typedef {import('./_components/types/user').Session} Session */

  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { IoClose } from './node_modules/svelte-icons-pack/dist/io';
  import { notifier } from './_components/xNotifier.js';
  import { datetime } from './_components/xformatter.js';
  import {  UserSessionKill } from './jsApi.GEN.js';
  import MainLayout from './_layouts/mainLayout.svelte';
  import ProfileSubMenu from './_components/partials/ProfileSubMenu.svelte';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});
  let sessions  = /** @type {Session[]} */ ([/* activeSessions */]);

  async function killSession(sessionToken) {
    await UserSessionKill({ sessionTokenHash: sessionToken }, async res => {
      // @ts-ignore
      if (res.error) return notifier.showError(res.error);
      if (res.sessionTerminated < 1) return notifier.showError('No session terminated');
      notifier.showInfo(res.sessionTerminated + ' session terminated');
      // @ts-ignore
      await UserSessionsActive({ userId: user.id }, res => {
        // @ts-ignore
        if (res.error) return notifier.showError(res.error);
        sessions = res.sessionsActive;
      });
    });
  }
</script>

<MainLayout>
  <ProfileSubMenu />
  <div class="sessions-container">
    <div class="header">
      <span>IP Address</span>
      <span>Expired At</span>
      <span>User Agent</span>
    </div>
    <div class="sessions">
      {#if sessions && sessions.length}
        {#each sessions as session}
          <div class="session">
            <span>{session.loginIPs || 'no-data'}</span>
            <span>{datetime(session.expiredAt) || 0}</span>
            <span>{session.device || 'no-data'}</span>
            <button
              on:click={() => killSession(session.sessionToken)}
              class="kill_session"
              title="kill this session"
            >
              <Icon color="#FFF" size="12" src={IoClose} />
            </button>
          </div>
        {/each}
      {:else}
        <div class="session">
          <span>No active sessions</span>
        </div>
      {/if}
    </div>
  </div>
</MainLayout>

<style>
  :global(.root-layout .root-container .root-content .content) {
    padding: 10px 20px 20px !important;
  }
  
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

  .sessions-container {
    border: 1px solid var(--gray-003);
    display: flex;
    flex-direction: column;
    height: fit-content;
    border-radius: 8px;
    overflow: hidden;
    background-color: #FFF;
  }

  .sessions-container .header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    font-weight: bold;
    padding: 12px 15px;
    border-bottom: 1px solid #cbd5e1;
  }

  .sessions-container .sessions {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .sessions-container .sessions .session {
    text-align: left;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 12px 15px;
    position: relative;
    width: 100%;
  }

  .sessions-container .sessions .session span:nth-child(3),
  .sessions-container .header span:nth-child(3) {
    flex-grow: 1;
    margin-right: 30px;
  }

  .sessions-container .sessions .session span,
  .sessions-container .header span {
    width: 200px;
  }

  .sessions-container .sessions .session .kill_session {
    border: none;
    background-color: var(--red-005);
    padding: 5px;
    border-radius: 50%;
    position: absolute;
    right: 15px;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .sessions-container .sessions .session .kill_session:hover {
    background-color: var(--red-004);
  }
</style>

<script>
  const year = new Date().getFullYear();
</script>

<footer>
  <p class="copyright">Copyright © {year}</p>
  <nav>
    <a href="/privacy">Privacy</a>
    <a href="/tos">Terms and Conditions </a>
    <a href="/about">About Us</a>
  </nav>
</footer>

<style>
  footer {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    margin: 0;
    padding: 20px 20px 0;
    border-top: 1px solid var(--gray-003);
    color: var(--gray-007);
  }
  footer p {
    margin: 0;
  }
  footer nav {
    display: flex;
    flex-direction: row;
    gap: 10px;
  }
  footer nav a {
    text-decoration: none;
    color: var(--gray-007);
  }
  footer nav a:hover {
    color: var(--sky-007);
    text-decoration: underline;
  }

  @media only screen and (max-width : 768px) {
    footer {
      flex-direction: column;
      gap: 10px;
    }
  }
</style>

import { expect, test } from '@playwright/test';

test('logs in with a valid username and password', async ({ page }) => {
  await page.goto('/login');

  await expect(page.getByTestId('login-form')).toBeVisible();

  await page.getByTestId('username').fill('demo@example.com');
  await page.getByTestId('password').fill('password');
  await page.getByTestId('login-submit').click();

  await expect(page).toHaveURL(/\/home$/);
  await expect(page.getByTestId('authenticated-banner')).toHaveText('You are signed in.');
  await expect(page.getByTestId('home-heading')).toContainText('Demo User');
});

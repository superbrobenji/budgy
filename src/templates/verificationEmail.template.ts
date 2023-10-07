export const emailSubject = (verificationCode: string) => {
  return `Verification Code: ${verificationCode}`; 
}
export const emailBody = (verificationCode: string) => {
  return `
    <h1>Welcome To Budgy!</h1>
    <p>Your verification code is: <b>${verificationCode}</b></p>
  `;
}

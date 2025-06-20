import { type TLoginForm, type TRegisterForm } from '@app/validation/auth';
import axiosInstance from '@app/lib/axios';
import type { User } from '@app/types/user';

export type ManualAuthResponse = {
  user: User;
  token: string;
};

export type AutoAuthResponse = {
  user: User;
};

export const getInitialize = () =>
  axiosInstance.get<AutoAuthResponse>('/auth/initialize');

export const postSignIn = (data: TLoginForm) =>
  axiosInstance.post<ManualAuthResponse>('/auth/sign-in', data);

export const postRegister = (data: TRegisterForm) =>
  axiosInstance.post<ManualAuthResponse>('/auth/register', {
    first_name: data.first_name,
    last_name: data.last_name,
    email: data.email,
    password: data.password,
    terms_and_conditions: true
  });

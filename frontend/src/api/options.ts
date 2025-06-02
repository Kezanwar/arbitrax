import axiosInstance from '@app/lib/axios';

interface ExchangesOptionsResp {
  exchanges: ExchangeItem[];
}

export interface ExchangeItem {
  key: string;
  label: string;
}

export const getExchangesOptions = () =>
  axiosInstance.get<ExchangesOptionsResp>('/options/exchanges');

interface StrategiesOptionsResp {
  strategies: StrategyItem[];
}

export interface StrategyItem {
  key: string;
  label: string;
  description: string;
}

export const getStrategiesOptions = () =>
  axiosInstance.get<StrategiesOptionsResp>('/options/strategies');

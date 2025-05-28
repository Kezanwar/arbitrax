import { getExchangesOptions, type ExchangeItem } from '@app/api/options';
import { useQuery } from '@tanstack/react-query';

export const EXCHANGES_QUERY_KEY = ['options', 'exchanges'];

export const useExchangesOptions = () => {
  return useQuery<ExchangeItem[]>({
    queryKey: EXCHANGES_QUERY_KEY,
    queryFn: async () => {
      const { data } = await getExchangesOptions();
      return data.exchanges;
    },
    staleTime: 1000 * 60 * 60 // 1 hour
  });
};

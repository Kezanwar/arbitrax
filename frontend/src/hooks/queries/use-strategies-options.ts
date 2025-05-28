import { getStrategiesOptions, type StrategyItem } from '@app/api/options';
import { useQuery } from '@tanstack/react-query';

export const STRATEGIES_QUERY_KEY = ['options', 'strategies'];

export const useStrategiesOptions = () => {
  return useQuery<StrategyItem[]>({
    queryKey: STRATEGIES_QUERY_KEY,
    queryFn: async () => {
      const { data } = await getStrategiesOptions();
      return data.strategies;
    },
    staleTime: 1000 * 60 * 60 // 1 hour
  });
};

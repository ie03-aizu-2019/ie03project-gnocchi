const endpoint = "http://localhost:5000";

export const EnumCrossPoints = "/enumCrossPoints";
export const RecomendCrossPoints = "/recomendCrossPoints";
export const DetectionHighWays = "/detectionHighWays";
export type ApiType =
  | typeof EnumCrossPoints
  | typeof RecomendCrossPoints
  | typeof DetectionHighWays;

export default (apiType: ApiType, payload?: RequestInit): Promise<string> =>
  fetch(endpoint + apiType, payload).then(x => x.text());

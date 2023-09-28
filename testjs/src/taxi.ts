export const Taxi = () => 55

export const fare = (distance: number, waitingTime: number): number => {
  return calculateFare(roundDistance(distance), roundWaitingTime(waitingTime))
}

export const calculateFare = (
  distance: number,
  waitingTime: number
): number => {
  const perKm = 2
  return perKm * distance + waitingTime
}

export const roundDistance = (distance: number): number => {
  return Math.ceil(distance * 2) / 2
}

export const roundWaitingTime = (waitingTime: number): number => {
  return Math.ceil(waitingTime)
}

import {
  Taxi,
  calculateFare,
  fare,
  roundDistance,
  roundWaitingTime,
} from './taxi'

describe('Taxi', () => {
  describe('calculateFare', () => {
    it('should return total of 1km without waiting time', () => {
      const distance = 1
      const waitingTime = 0

      const result = calculateFare(distance, waitingTime)

      expect(result).toEqual(2)
    })

    it('should return total of 2km without waiting time', () => {
      const distance = 2
      const waitingTime = 0

      const result = calculateFare(distance, waitingTime)

      expect(result).toEqual(4)
    })

    it('should return total of 0km with 1m waiting time', () => {
      const distance = 0
      const waitingTime = 1

      const result = calculateFare(distance, waitingTime)

      expect(result).toEqual(1)
    })

    it('should return total of 2km with 2m waiting time', () => {
      const distance = 2
      const waitingTime = 2

      const result = calculateFare(distance, waitingTime)

      expect(result).toEqual(6)
    })
  })

  describe('rounding distance', () => {
    it('should round up distance to 0.5km', () => {
      const distance = 0.1

      const result = roundDistance(distance)

      expect(result).toEqual(0.5)
    })

    it('should round up distance to 0.5km', () => {
      const distance = 0.5

      const result = roundDistance(distance)

      expect(result).toEqual(0.5)
    })

    it('should round up distance to 0.6km', () => {
      const distance = 0.6

      const result = roundDistance(distance)

      expect(result).toEqual(1)
    })

    it('should round up distance to 1.0km', () => {
      const distance = 1.0

      const result = roundDistance(distance)

      expect(result).toEqual(1.0)
    })

    it('should round up distance to 1.3km', () => {
      const distance = 1.3

      const result = roundDistance(distance)

      expect(result).toEqual(1.5)
    })
  })

  describe('rounding waitingTime', () => {
    it('should round up waitingTime to 1m', () => {
      const waitingTime = 0.1

      const result = roundWaitingTime(waitingTime)

      expect(result).toEqual(1)
    })

    it('should round up waitingTime to 2m', () => {
      const waitingTime = 1.2

      const result = roundWaitingTime(waitingTime)

      expect(result).toEqual(2)
    })

    it('should round up waitingTime to 2m', () => {
      const waitingTime = 1.5

      const result = roundWaitingTime(waitingTime)

      expect(result).toEqual(2)
    })
  })

  describe('fare of ride', () => {
    it('should return cost of 1.3km without waiting time', () => {
      const distance = 1.3
      const waitingTime = 0

      const result = fare(distance, waitingTime)

      expect(result).toEqual(3)
    })

    it('should return cost of 1.5km without waiting time', () => {
      const distance = 1.5
      const waitingTime = 0

      const result = fare(distance, waitingTime)

      expect(result).toEqual(3)
    })

    it('should return cost of 2km without waiting time', () => {
      const distance = 2
      const waitingTime = 0

      const result = fare(distance, waitingTime)

      expect(result).toEqual(4)
    })
    it('should return cost of 0km with 1m waiting time', () => {
      const distance = 0
      const waitingTime = 1

      const result = fare(distance, waitingTime)

      expect(result).toEqual(1)
    })

    it('should return cost of 0km with 1.2m waiting time', () => {
      const distance = 0
      const waitingTime = 1.2

      const result = fare(distance, waitingTime)

      expect(result).toEqual(2)
    })

    it('should return cost of 0km with 1.5m waiting time', () => {
      const distance = 0
      const waitingTime = 1.5

      const result = fare(distance, waitingTime)

      expect(result).toEqual(2)
    })

    it('should return cost of 0km with 1.6m waiting time', () => {
      const distance = 0
      const waitingTime = 1.6

      const result = fare(distance, waitingTime)

      expect(result).toEqual(2)
    })

    it('should return cost of 8.2km with 2.6m waiting time', () => {
      const distance = 8.2
      const waitingTime = 2.6

      const result = fare(distance, waitingTime)

      expect(result).toEqual(20)
    })
  })
})

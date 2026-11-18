"""Loyalty points for orders: a new service, not in the model yet."""

POINTS_PER_UNIT = 10


def points_for(units):
    return max(0, int(units)) * POINTS_PER_UNIT

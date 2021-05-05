def powersum(data):
    sum = 0.0
    for i, v in enumerate(data):
        sum += abs(v[0] + v[1])

    return sum/data.shape[0]

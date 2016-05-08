#!/usr/bin/env python3

import itertools

T = int(input())

digits = "0123456789"

def solve(J, P, S, K):
    jp = {}
    js = {}
    ps = {}
    jps = {}
    for j in range(J):
        for p in range(P):
            jp[str(j) + str(p)] = 0
            for s in range(S):
                jps[str(j) + str(p) + str(s)] = False
        for s in range(S):
            js[str(j) + str(s)] = 0
    for p in range(P):
        for s in range(S):
            ps[str(p) + str(s)] = 0

    res = []
    added = True

    while added:
        added = False
        for j, p, s in itertools.product(digits[:J], digits[:P], digits[:S]):
            if jp[j + p] < K and  js[j + s] < K and ps[p + s] < K and not jps[j + p + s]:
                jp[j + p] += 1
                js[j + s] += 1
                ps[p + s] += 1
                jps[j + p + s] = True
                added = True
                res.append([str(int(n) + 1) for n in [j, p, s]])

    return res

for t in range(0, T):
    J, P, S, K = [int(n) for n in input().split(" ")]
    solutions = solve(J, P, S, K)
    print("Case #%d: %d" % (t + 1, len(solutions)))
    for solution in solutions:
        print(" ".join(solution))

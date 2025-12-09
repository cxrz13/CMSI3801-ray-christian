using Distributed
addprocs(parse(Int, ARGS[1]))

@everywhere function approximate_pi(trials::Int)
    hits = 0
    for i in 1:trials
        hits += (rand()^2 + rand()^2 < 1) ? 1 : 0
    end
    return hits
end

function main()
    total_trials = 500_000_000
    trials_per_worker = div(total_trials, nworkers())
    hits = pmap(w -> approximate_pi(trials_per_worker), workers())
    return 4 * sum(hits) / total_trials
end

println("Estimating π with $(nworkers()) workers...")
@time estimate = main()
println("π ≈ $estimate")

# Logs:
# 1 worker: time = 2.433567 seconds
# 2 workers: time = 1.939847 seconds
# 3 workers: time =  1.772266 seconds
# 4 workers: time = 1.631783 seconds
# 5 workers: time = 1.702577 seconds
# 6 workers: time = 2.160641 seconds
# 7 workers: time = 2.000053 seconds
# 8 workers: time =  2.018001 seconds
# 9 workers: time =  2.281580 seconds
# 10 workers: time = 3.485231 seconds
# 11 workers: time = 3.928171 seconds
# 12 workers: time = 4.322148 seconds
# 13 workers: time = 4.838249 seconds
# 14 workers: time = 5.404653 seconds
# 15 workers: time = 6.144642 seconds
# 16 workers: time = 6.984643 seconds
# 17 workers: time = 6.695797 seconds
# 18 workers: time = 7.246139 seconds
# 19 workers: time = 9.516953 seconds
# 20 workers: time = 9.874331 seconds
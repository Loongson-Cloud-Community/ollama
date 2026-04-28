cmake --preset 'CPU' \
        && cmake --build --preset 'CPU' -- -l $(nproc) \
        && cmake --install build --component CPU --strip


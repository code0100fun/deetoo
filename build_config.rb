def gem_config(conf)
  conf.gem :git => 'https://github.com/matsumoto-r/mruby-sleep.git'
  conf.gembox 'default'
end

MRuby::Build.new do |conf|
  toolchain :gcc
  enable_debug

  gem_config(conf)
end

MRuby::Build.new('host-debug') do |conf|
  toolchain :gcc

  enable_debug

  conf.cc.defines = %w(MRB_ENABLE_DEBUG_HOOK)
  conf.gem :core => "mruby-bin-debugger"

  gem_config(conf)
end

MRuby::Build.new('test') do |conf|
  toolchain :gcc

  enable_debug
  conf.enable_bintest
  conf.enable_test

  gem_config(conf)
end
